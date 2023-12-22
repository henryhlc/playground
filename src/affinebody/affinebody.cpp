#include "./affinebody.h"

#include <iomanip>

#include <ipc/distance/distance_type.hpp>
#include <ipc/ccd/ccd.hpp>

#include "./contact_ip.h"
#include "./orthogonal_ip.h"
#include "./massmatrix_integral.h"

constexpr double kappa = 0.9;
constexpr double stiffness = 200.0;
constexpr int max_newton_iter = 15;
constexpr int max_line_iter = 8;

Eigen::Matrix<double,12,12> mass_matrix(
        const Eigen::MatrixX3d& V,
        const Eigen::MatrixX3i& F) {

    Eigen::Matrix<double,12,12> M = Eigen::Matrix<double,12,12>::Zero();

    for (int i = 0; i < F.rows(); ++i) {
        auto v1 = V.row(F(i,0));
        auto x1 = v1(0);
        auto y1 = v1(1);
        auto z1 = v1(2);
        auto v2 = V.row(F(i,1));
        auto x2 = v2(0);
        auto y2 = v2(1);
        auto z2 = v2(2);
        auto v3 = V.row(F(i,2));
        auto x3 = v3(0);
        auto y3 = v3(1);
        auto z3 = v3(2);

        auto m_1 = I_1(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_x = I_x(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_y = I_y(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_z = I_z(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_xx = I_xx(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_yy = I_yy(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_zz = I_zz(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_xy = I_xy(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_yz = I_yz(x1, y1, z1, x2, y2, z2, x3, y3, z3);
        auto m_xz = I_xz(x1, y1, z1, x2, y2, z2, x3, y3, z3);

        Eigen::Matrix3d top_left;
        top_left << m_1, 0, 0,
                    0, m_1, 0,
                    0, 0, m_1;
        M.block<3,3>(0,0) += top_left;
        Eigen::Matrix3d cross_block;
        cross_block << m_xx, m_xy, m_xz,
                       m_xy, m_yy, m_yz,
                       m_xz, m_yz, m_zz;
        M.block<3,3>(3,3) += cross_block;
        M.block<3,3>(6,6) += cross_block;
        M.block<3,3>(9,9) += cross_block;

        Eigen::Vector3d xyz;
        xyz << m_x, m_y, m_z;
        M.block<3,1>(3,0) += xyz;
        M.block<3,1>(6,1) += xyz;
        M.block<3,1>(9,2) += xyz;
        M.block<1,3>(0,3) += xyz;
        M.block<1,3>(1,6) += xyz;
        M.block<1,3>(2,9) += xyz;
    }
    return M;
}

Eigen::VectorXd stack_affine_body_coordinates(
    const std::vector<AffineBody>& abs) {

    Eigen::VectorXd q {12 * abs.size()};
    for (int i = 0; i < abs.size(); ++i) {
        auto& ab = abs[i];
        q(Eigen::seqN(12 * i, 12)) = ab.q;
    }
    return q;
}

std::vector<AffineBody> unstack_affine_body_coordinates(
    Eigen::Ref<Eigen::VectorXd> q,
    const std::vector<AffineBody>& abs) {

    std::vector<AffineBody> states; 
    for (int i = 0; i < abs.size(); ++i) {
        auto& ab = abs[i];
        states.emplace_back(ab.world_V, ab.F, q(Eigen::seqN(i*12, 12)));
    }
    return states;
}

Eigen::MatrixXd project_PSD(Eigen::MatrixXd X) {
    Eigen::SelfAdjointEigenSolver<Eigen::MatrixXd> es{X};
    Eigen::MatrixXd D = es.eigenvalues().cwiseMax(0.05).asDiagonal();
    Eigen::MatrixXd V = es.eigenvectors();
    return V * D * V.inverse();
}

double time_of_impact_pt_pe(
    const AffineBody& abi_before,
    const AffineBody& abj_before,
    const AffineBody& abi_after,
    const AffineBody& abj_after) {
    auto Vi_before = abi_before.V();
    auto Vi_after = abi_after.V();
    auto Vj_before = abj_before.V();
    auto Vj_after = abj_after.V();
    auto Fj = abj_before.F;

    double toi = 1.0;

    for (int vi_i = 0; vi_i < Vi_before.rows(); ++vi_i) {
        for (int fj_i = 0; fj_i < Fj.rows(); ++fj_i) {
            auto p_before = Vi_before.row(vi_i);
            auto p1_before = Vj_before.row(Fj(fj_i,0));
            auto p2_before = Vj_before.row(Fj(fj_i,1));
            auto p3_before = Vj_before.row(Fj(fj_i,2));
            auto p_after = Vi_after.row(vi_i);
            auto p1_after = Vj_after.row(Fj(fj_i,0));
            auto p2_after = Vj_after.row(Fj(fj_i,1));
            auto p3_after = Vj_after.row(Fj(fj_i,2));

            double toi_pt;
            ipc::point_triangle_ccd(p_before,p1_before,p2_before,p3_before,p_after,p1_after,p2_after,p3_after,toi_pt);
            toi = std::min(toi,toi_pt);

            for (int e = 0; e < 3; ++e) {
                auto e1_before = Vj_before.row(Fj(fj_i,e));
                auto e2_before = Vj_before.row(Fj(fj_i,(e+1)%3));
                auto e1_after = Vj_after.row(Fj(fj_i,e));
                auto e2_after = Vj_after.row(Fj(fj_i,(e+1)%3));

                double toi_pe;
                ipc::point_edge_ccd_3D(p_before,e1_before,e2_before,p_after,e1_after,e2_after,toi_pe);
                toi = std::min(toi,toi_pe);
            }
        }
    }
    return toi;
}

double time_of_impact(
    const std::vector<AffineBody>& before,
    const std::vector<AffineBody>& after) {

    double toi = 1.0;

    // pp, ee, pt, pe
    for (int b_i = 1; b_i < before.size(); ++b_i) {
        for (int b_j = 0; b_j < b_i; ++b_j) {
            auto& abi_before = before[b_i];
            auto& abi_after = after[b_i];
            auto Vi_before = abi_before.V();
            auto Vi_after = abi_after.V();
            auto& Fi = abi_before.F;

            auto& abj_before = before[b_j];
            auto& abj_after = after[b_j];
            auto Vj_before = abj_before.V();
            auto Vj_after = abj_after.V();
            auto& Fj = abj_before.F;

            toi = std::min(toi, time_of_impact_pt_pe(abi_before,abj_before,abi_after,abj_after));
            toi = std::min(toi, time_of_impact_pt_pe(abj_before,abi_before,abj_after,abi_after));

            for (int vi_i = 0; vi_i < Vi_before.rows(); ++vi_i) {
                for (int vj_i = 0; vj_i < Vj_before.rows(); ++vj_i) {
                    auto p1_before = Vi_before.row(vi_i);
                    auto p1_after = Vi_after.row(vi_i);
                    auto p2_before = Vj_before.row(vj_i);
                    auto p2_after = Vj_after.row(vj_i);
                    double toi_pp;
                    ipc::point_point_ccd_3D(p1_before,p2_before,p1_after,p2_after,toi_pp);
                    toi = std::min(toi,toi_pp);
                }
            }

            for (int fi_i = 0; fi_i < Fi.rows(); ++fi_i) {
                for (int fj_i = 0; fj_i < Fj.rows(); ++fj_i) {
                    for (int e1_start = 0; e1_start < 3; ++e1_start) {
                        for (int e2_start = 0; e2_start < 3; ++e2_start) {
                            auto p11_before = Vi_before.row(Fi(fi_i,e1_start));
                            auto p12_before = Vi_before.row(Fi(fi_i,(e1_start+1)%3));
                            auto p21_before = Vj_before.row(Fj(fj_i,e2_start));
                            auto p22_before = Vj_before.row(Fj(fj_i,(e2_start+1)%3));
                            auto p11_after = Vi_after.row(Fi(fi_i,e1_start));
                            auto p12_after = Vi_after.row(Fi(fi_i,(e1_start+1)%3));
                            auto p21_after = Vj_after.row(Fj(fj_i,e2_start));
                            auto p22_after = Vj_after.row(Fj(fj_i,(e2_start+1)%3));
                            double toi_ee;
                            ipc::edge_edge_ccd(p11_before,p12_before,p21_before,p22_before,
                                p11_after,p12_after,p21_after,p22_after,toi_ee);
                            toi = std::min(toi,toi_ee);
                        }
                    }
                }
            }
        }
    }
    return toi;
}

AffineBody::AffineBody(
    const Eigen::MatrixX3d world_V,
    const Eigen::MatrixX3i F,
    Eigen::Vector<double,12> q
): world_V{world_V}, F{F}, q{q}, M{mass_matrix(world_V,F)} {}

double incremental_potential(
        Eigen::Ref<Eigen::VectorXd> q,
        const std::vector<AffineBody>& curr_states,
        const std::vector<AffineBody>& prev_states,
        std::chrono::milliseconds dt) {
    
    double dt_s = std::chrono::duration<double>(dt).count();

    double ip = 0;

    std::cout << std::fixed << std::setprecision(6);

    for (int i = 0; i < curr_states.size(); ++i) {
        auto& ab = curr_states[i];
        auto& ab_prev = prev_states[i];
        
        auto q_gap = q(Eigen::seqN(i*12,12)) - (2*ab.q - ab_prev.q); // + dt^2 M^{-1}F
        double ip_k = 0.5 * q_gap.transpose() * ab.M * q_gap;
        ip += ip_k;
        std::cout << "kinetic: " << ip_k << "; ";

        double ip_o = stiffness * dt_s * dt_s * ip_orthogonal(q(Eigen::seqN(i*12,12)));
        ip += ip_o;
        std::cout << "orthogonal: " << ip_o << "; ";
    }

    auto q_states = unstack_affine_body_coordinates(q,curr_states);
    double ip_contact = kappa * contact_ip(q, q_states);
    ip += ip_contact;
    std::cout << "contact: " << ip_contact << "; ";

    // TODO: friction term

    std::cout << "total: " << ip << std::endl;
    return ip;
}

std::vector<AffineBody> affine_body_dynamics(
        const std::vector<AffineBody>& curr_states,
        const std::vector<AffineBody>& prev_states,
        std::chrono::milliseconds dt) {

    double dt_s = std::chrono::duration<double>(dt).count();

    auto q_curr = stack_affine_body_coordinates(curr_states);
    auto q_prev = stack_affine_body_coordinates(prev_states);
    auto q_free = 2 * q_curr - q_prev;  // + dt^2 M^{-1}F

    // std::cout << "current" << std::endl;
    // std::cout << q_curr.transpose() << std::endl;
    // std::cout << "free" << std::endl;
    // std::cout << q_free.transpose() << std::endl;

    Eigen::VectorXd q_iter = q_curr;
    double ip_iter = incremental_potential(q_iter, curr_states, prev_states, dt);

    int num_newton_iter = 0;
    while (true) {
        if (num_newton_iter >= max_newton_iter) {
            std::cout << "Newton iterations has not converged after " << max_newton_iter << " iterations" << std::endl;
            break;
        }
        num_newton_iter++;
        // std::cout << std::endl << "iter" << std::endl;
        // std::cout << q_iter.transpose() << std::endl;

        auto q_iter_states = unstack_affine_body_coordinates(q_iter, curr_states);

        // Search direction components
        Eigen::VectorXd grad = Eigen::VectorXd::Zero(12*q_iter_states.size());
        Eigen::MatrixXd hess = Eigen::MatrixXd::Zero(12*q_iter_states.size(), 12*q_iter_states.size());

        // Per body incremental potential gradients and hessian components
        for (int i = 0; i < q_iter_states.size(); ++i) {
            auto& ab = q_iter_states[i];
        
            // Kinetic term
            auto q_gap = ab.q - q_free(Eigen::seqN(i*12,12));
            auto half_M_MT = 0.5 * (ab.M + ab.M.transpose());
            grad(Eigen::seqN(i*12,12)) += half_M_MT * q_gap;
            hess.block<12,12>(i*12,i*12) += half_M_MT;

            // Stiffness term
            grad(Eigen::seqN(i*12,12)) += stiffness * dt_s * dt_s * grad_orthogonal(ab.q);
            hess.block<12,12>(i*12,i*12) += stiffness * dt_s * dt_s * project_PSD(hess_orthogonal(ab.q));
        }

        // Contact term
        grad += kappa * contact_ip_gradient(q_iter, q_iter_states);
        hess += kappa * project_PSD(contact_ip_hessian(q_iter, q_iter_states));

        // TODO: friction term

        Eigen::VectorXd search_direction = -project_PSD(hess).fullPivLu().solve(grad);
        // std::cout << "direction" << std::endl;
        // std::cout << search_direction.transpose() << std::endl;

        double search_direction_inf_norm = search_direction.lpNorm<Eigen::Infinity>();
        if (search_direction_inf_norm < 1e-3) {
            break;
        }

        // Line search

        Eigen::VectorXd q_max_step = q_iter + search_direction;
        auto q_max_step_states = unstack_affine_body_coordinates(q_max_step, q_iter_states);
        double max_step_size = std::min(1.0, time_of_impact(q_iter_states,q_max_step_states));

        double step_size = max_step_size;

        int num_line_iter = 0;

        Eigen::VectorXd q_cand;
        double ip_cand;
        while (true) {
            if (num_line_iter >= max_line_iter) {
                std::cout << "Line search target not found after " << max_line_iter << " iterations" << std::endl;
                break;
            }
            num_line_iter++;

            q_cand = q_iter + step_size * search_direction;
            // std::cout << "candidate" << std::endl;
            // std::cout << q_cand.transpose() << std::endl;
            ip_cand = incremental_potential(q_cand, curr_states, prev_states, dt);


            if (ip_cand < ip_iter) {
                break;
            } else {
                step_size /= 2;
            }
        }
        q_iter = q_cand;
        ip_iter = ip_cand;
    } 

    return unstack_affine_body_coordinates(q_iter, curr_states);
}
