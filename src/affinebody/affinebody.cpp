#include "affinebody.h"

#include "./massmatrix_integral.h"

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

AffineBody::AffineBody(
    const Eigen::Ref<const Eigen::MatrixX3d> world_V,
    const Eigen::Ref<const Eigen::MatrixX3i> F,
    Eigen::Vector<double,12> q
): world_V{world_V}, F{F}, q{q}, M{mass_matrix(world_V,F)} {}

std::vector<AffineBody> affine_body_dynamics(
        const std::vector<AffineBody>& curr_states,
        const std::vector<AffineBody>& prev_states,
        std::chrono::milliseconds dt) {

    Eigen::VectorXd q_curr {12 * curr_states.size()};
    Eigen::VectorXd q_prev {12 * prev_states.size()};
    for (int i = 0; i < curr_states.size(); ++i) {
        auto& ab = curr_states[i];
        q_curr(Eigen::seqN(12 * i, 12)) = ab.q;
        auto& ab_prev = prev_states[i];
        q_prev(Eigen::seqN(12 * i, 12)) = ab_prev.q;

    }

    Eigen::VectorXd q_iter = q_curr;

    int num_newton_iter = 0;
    int max_newton_iter = 1;
    while (true) {
        if (num_newton_iter >= max_newton_iter) {
            std::cout << "Newton iterations has not converged after " << max_newton_iter << " iterations" << std::endl;
            break;
        }
        num_newton_iter++;

        // direction
        double ie = 0.0;
        Eigen::VectorXd grad = Eigen::VectorXd::Zero(12*curr_states.size());
        Eigen::MatrixXd hess = Eigen::MatrixXd::Zero(12*curr_states.size(), 12*curr_states.size());

        for (int i = 0; i < curr_states.size(); ++i) {
            auto& ab = curr_states[i];
            auto& ab_prev = prev_states[i];
        
            // Kinetic term

            auto q_gap = q_iter(Eigen::seqN(i*12,12)) - (2*ab.q - ab_prev.q); // + dt^2 M^{-1}F
            auto half_M_MT = 0.5 * ab.M + ab.M.transpose();

            ie += 0.5 * q_gap.transpose() * ab.M * q_gap;
            grad(Eigen::seqN(i*12, 12)) += half_M_MT * q_gap;
            hess.block<12,12>(i*12, i*12) += half_M_MT;

            // TODO: Orthogonal term
        }

        // TODO: contact
        // all-pairs of body
        // all-pairs vertex-face
        // all-pairs edge-edge

        // TODO: friction

        Eigen::VectorXd search_direction = -hess.fullPivLu().solve(grad);

        if (search_direction.lpNorm<Eigen::Infinity>() < 1e-4) {
            break;
        }

        // Line search

        // compute max with CCD
        double max_step_size = 1;
        double step_size = max_step_size;

        int num_line_iter = 0;
        int max_line_iter = 1;
        while (true) {
            if (num_line_iter >= max_line_iter) {
                std::cout << "Line search target not found after " << max_line_iter << " iterations" << std::endl;
                break;
            }
            num_line_iter++;

            auto q_cand = q_iter + step_size * search_direction;

            // TODO: compute energy of candidate
            double ie_cand = 0.0;

            if (ie_cand < ie) {
                q_iter = q_cand;
                break;
            } else {
                step_size /= 2;
            }
        }
    } 

    std::vector<AffineBody> next_states; 
    for (int i = 0; i < curr_states.size(); ++i) {
        auto& ab_curr = curr_states[i];
        next_states.emplace_back(ab_curr.world_V, ab_curr.F, q_iter(Eigen::seqN(i*12, 12)));
    }

    return next_states;
}
