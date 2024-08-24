#include "contact_ip.h"

#include <functional>

#include <ipc/distance/distance_type.hpp>

#include "./distance.h"

// TODO: include affine body, and use the impl there.
Eigen::Vector3d from_body_space(Eigen::Vector<double,12> q, Eigen::Vector3d p) {
    auto d = q(Eigen::seqN(0,3));
    Eigen::Matrix3d A;
    A.row(0) = q(Eigen::seqN(3,3));
    A.row(1) = q(Eigen::seqN(6,3));
    A.row(2) = q(Eigen::seqN(9,3));
    return d + A * p;
}

double d2_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    double d = 0.0;
    switch (ipc::point_triangle_distance_type(
        from_body_space(q1,p),
        from_body_space(q2,p1),
        from_body_space(q2,p2),
        from_body_space(q2,p3))) {
        case ipc::PointTriangleDistanceType::P_T0:
            d = d2_pp(q1,p,q2,p1);
            break;
        case ipc::PointTriangleDistanceType::P_T1:
            d = d2_pp(q1,p,q2,p2);
            break;
        case ipc::PointTriangleDistanceType::P_T2:
            d = d2_pp(q1,p,q2,p3);
            break;
        case ipc::PointTriangleDistanceType::P_E0:
            d = d2_pe(q1,p,q2,p1,p2);
            break;
        case ipc::PointTriangleDistanceType::P_E1:
            d = d2_pe(q1,p,q2,p2,p3);
            break;
        case ipc::PointTriangleDistanceType::P_E2:
            d = d2_pe(q1,p,q2,p3,p1);
            break;
        case ipc::PointTriangleDistanceType::P_T:
            d = d2_pt(q1,p,q2,p1,p2,p3);
    }
    return d;
}

double ip_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    double d {d2_vertex_face(q1,p,q2,p1,p2,p3)};
    return barrier(d);
}

Eigen::Vector<double,24> gradient_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    Eigen::Vector<double,24> d_grad;
    switch (ipc::point_triangle_distance_type(
        from_body_space(q1,p),
        from_body_space(q2,p1),
        from_body_space(q2,p2),
        from_body_space(q2,p3))) {
        case ipc::PointTriangleDistanceType::P_T0:
            d_grad = d2_pp_grad(q1,p,q2,p1);
            break;
        case ipc::PointTriangleDistanceType::P_T1:
            d_grad = d2_pp_grad(q1,p,q2,p2);
            break;
        case ipc::PointTriangleDistanceType::P_T2:
            d_grad = d2_pp_grad(q1,p,q2,p3);
            break;
        case ipc::PointTriangleDistanceType::P_E0:
            d_grad = d2_pe_grad(q1,p,q2,p1,p2);
            break;
        case ipc::PointTriangleDistanceType::P_E1:
            d_grad = d2_pe_grad(q1,p,q2,p2,p3);
            break;
        case ipc::PointTriangleDistanceType::P_E2:
            d_grad = d2_pe_grad(q1,p,q2,p3,p1);
            break;
        case ipc::PointTriangleDistanceType::P_T:
            d_grad = d2_pt_grad(q1,p,q2,p1,p2,p3);
            break;
    }
    double d {d2_vertex_face(q1,p,q2,p1,p2,p3)};
    return d_barrier(d) * d_grad;
}

Eigen::Matrix<double,24,24> hessian_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    Eigen::Matrix<double,24,24> d_hess;
    switch (ipc::point_triangle_distance_type(
        from_body_space(q1,p),
        from_body_space(q2,p1),
        from_body_space(q2,p2),
        from_body_space(q2,p3))) {
        case ipc::PointTriangleDistanceType::P_T0:
            d_hess = d2_pp_hess(q1,p,q2,p1);
            break;
        case ipc::PointTriangleDistanceType::P_T1:
            d_hess = d2_pp_hess(q1,p,q2,p2);
            break;
        case ipc::PointTriangleDistanceType::P_T2:
            d_hess = d2_pp_hess(q1,p,q2,p3);
            break;
        case ipc::PointTriangleDistanceType::P_E0:
            d_hess = d2_pe_hess(q1,p,q2,p1,p2);
            break;
        case ipc::PointTriangleDistanceType::P_E1:
            d_hess = d2_pe_hess(q1,p,q2,p2,p3);
            break;
        case ipc::PointTriangleDistanceType::P_E2:
            d_hess = d2_pe_hess(q1,p,q2,p3,p1);
            break;
        case ipc::PointTriangleDistanceType::P_T:
            d_hess = d2_pt_hess(q1,p,q2,p1,p2,p3);
            break;
    }
    double d {d2_vertex_face(q1,p,q2,p1,p2,p3)};
    Eigen::Vector<double,24> d_grad = gradient_vertex_face(q1,p,q2,p1,p2,p3);
    return d_barrier(d) * d_hess + dd_barrier(d) * d_grad * d_grad.transpose();
}

double d2_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {
    switch (ipc::edge_edge_distance_type(
        from_body_space(q1,p11),
        from_body_space(q1,p12),
        from_body_space(q2,p21),
        from_body_space(q2,p22))) {
        case ipc::EdgeEdgeDistanceType::EA0_EB0:
            return d2_pp(q1,p11,q2,p21);
        case ipc::EdgeEdgeDistanceType::EA0_EB1:
            return d2_pp(q1,p11,q2,p22);
        case ipc::EdgeEdgeDistanceType::EA1_EB0:
            return d2_pp(q1,p12,q2,p21);
        case ipc::EdgeEdgeDistanceType::EA1_EB1:
            return d2_pp(q1,p12,q2,p22);
        case ipc::EdgeEdgeDistanceType::EA_EB0:
            return d2_pe(q2,p21,q1,p11,p12);
        case ipc::EdgeEdgeDistanceType::EA_EB1:
            return d2_pe(q2,p22,q1,p11,p12);
        case ipc::EdgeEdgeDistanceType::EA0_EB:
            return d2_pe(q1,p11,q2,p21,p22);
        case ipc::EdgeEdgeDistanceType::EA1_EB:
            return d2_pe(q1,p12,q2,p21,p22);
        case ipc::EdgeEdgeDistanceType::EA_EB:
            return d2_ee(q1,p11,p12,q2,p21,p22);
    }
    return 0.0;
}

double ip_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {
    double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
    return barrier(d);
}

Eigen::Vector<double,24> gradient_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {

    Eigen::Vector<double,24> d_grad;
    switch (ipc::edge_edge_distance_type(
        from_body_space(q1,p11),
        from_body_space(q1,p12),
        from_body_space(q2,p21),
        from_body_space(q2,p22))) {
        case ipc::EdgeEdgeDistanceType::EA0_EB0:
            d_grad = d2_pp_grad(q1,p11,q2,p21);
            break;
        case ipc::EdgeEdgeDistanceType::EA0_EB1:
            d_grad = d2_pp_grad(q1,p11,q2,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA1_EB0:
            d_grad = d2_pp_grad(q1,p12,q2,p21);
            break;
        case ipc::EdgeEdgeDistanceType::EA1_EB1:
            d_grad = d2_pp_grad(q1,p12,q2,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA_EB0: {
            Eigen::Vector<double,24> grad;
            auto pe_grad = d2_pe_grad(q2,p21,q1,p11,p12);
            grad(Eigen::seqN(0,12)) = pe_grad(Eigen::seqN(12,12));
            grad(Eigen::seqN(12,12)) = pe_grad(Eigen::seqN(0,12));
            d_grad = grad;
            break;
        }
        case ipc::EdgeEdgeDistanceType::EA_EB1: {
            Eigen::Vector<double,24> grad;
            auto pe_grad = d2_pe_grad(q2,p22,q1,p11,p12);
            grad(Eigen::seqN(0,12)) = pe_grad(Eigen::seqN(12,12));
            grad(Eigen::seqN(12,12)) = pe_grad(Eigen::seqN(0,12));
            d_grad = grad;
            break;
        }
        case ipc::EdgeEdgeDistanceType::EA0_EB:
            d_grad = d2_pe_grad(q1,p11,q2,p21,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA1_EB:
            d_grad = d2_pe_grad(q1,p12,q2,p21,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA_EB:
            d_grad = d2_ee_grad(q1,p11,p12,q2,p21,p22);
            break;
    }
    double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
    return d_barrier(d) * d_grad;
}

Eigen::Matrix<double,24,24> hessian_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {

    Eigen::Matrix<double,24,24> d_hess;
    switch (ipc::edge_edge_distance_type(
        from_body_space(q1,p11),
        from_body_space(q1,p12),
        from_body_space(q2,p21),
        from_body_space(q2,p22))) {
        case ipc::EdgeEdgeDistanceType::EA0_EB0:
            d_hess = d2_pp_hess(q1,p11,q2,p21);
            break;
        case ipc::EdgeEdgeDistanceType::EA0_EB1:
            d_hess = d2_pp_hess(q1,p11,q2,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA1_EB0:
            d_hess = d2_pp_hess(q1,p12,q2,p21);
            break;
        case ipc::EdgeEdgeDistanceType::EA1_EB1:
            d_hess = d2_pp_hess(q1,p12,q2,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA_EB0: {
            Eigen::Matrix<double,24,24> hess;
            auto pe_hess = d2_pe_hess(q2,p21,q1,p11,p12);
            hess.block<12,12>(0,0) = pe_hess.block<12,12>(12,12);
            hess.block<12,12>(12,0) = pe_hess.block<12,12>(0,12);
            hess.block<12,12>(0,12) = pe_hess.block<12,12>(12,0);
            hess.block<12,12>(12,12) = pe_hess.block<12,12>(0,0);
            d_hess = hess;
            break;
        }
        case ipc::EdgeEdgeDistanceType::EA_EB1: {
            Eigen::Matrix<double,24,24> hess;
            auto pe_hess = d2_pe_hess(q2,p22,q1,p11,p12);
            hess.block<12,12>(0,0) = pe_hess.block<12,12>(12,12);
            hess.block<12,12>(12,0) = pe_hess.block<12,12>(0,12);
            hess.block<12,12>(0,12) = pe_hess.block<12,12>(12,0);
            hess.block<12,12>(12,12) = pe_hess.block<12,12>(0,0);
            d_hess = hess;
            break;
        }
        case ipc::EdgeEdgeDistanceType::EA0_EB:
            d_hess = d2_pe_hess(q1,p11,q2,p21,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA1_EB:
            d_hess = d2_pe_hess(q1,p12,q2,p21,p22);
            break;
        case ipc::EdgeEdgeDistanceType::EA_EB:
            d_hess = d2_ee_hess(q1,p11,p12,q2,p21,p22);
            break;
    }
    double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
    Eigen::Vector<double,24> d_grad = gradient_edge_edge(q1,p11,p12,q2,p21,p22);
    return d_barrier(d) * d_hess + dd_barrier(d) * d_grad * d_grad.transpose();
}

using VertexFaceFn = std::function<void(
    int i,
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    int j,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3)>;

using FaceVertexFn = std::function<void(
    int i,
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3,
    int j,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p)>;

using EdgeEdgeFn = std::function<void(
    int i,
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    int j,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22)>;


void iterate_primitives(
    Eigen::VectorXd q,
    std::vector<AffineBody> abs,
    VertexFaceFn vertex_face_fn,
    FaceVertexFn face_vertex_fn,
    EdgeEdgeFn edge_edge_fn) {

    // all pairs of bodies
    for (int i = 1; i < abs.size(); ++i) {
        for (int j = 0; j < i; ++j) {
            auto& abi = abs[i];
            auto& abj = abs[j];

            // all pairs of faces
            for (auto fi : abi.F.rowwise()) {
                for (auto fj : abj.F.rowwise()) {
                    auto vi1 = abi.world_V.row(fi[0]);
                    auto vi2 = abi.world_V.row(fi[1]);
                    auto vi3 = abi.world_V.row(fi[2]);
                    auto vj1 = abj.world_V.row(fj[0]);
                    auto vj2 = abj.world_V.row(fj[1]);
                    auto vj3 = abj.world_V.row(fj[2]);

                    // all face-vertex
                    for (auto vj_index : fj) {
                        // fi and vj
                        Eigen::Vector3d vj = abj.world_V.row(vj_index);
                        face_vertex_fn(i,abi.q,vi1,vi2,vi3,j,abj.q,vj);
                    }
                    // all vertex-face
                    for (auto vi_index : fi) {
                        // vi and fj
                        Eigen::Vector3d vi = abi.world_V.row(vi_index);
                        vertex_face_fn(i,abi.q,vi,j,abj.q,vj1,vj2,vj3);

                    }
                    // all edge-edge
                    for (int ei_start = 0; ei_start < 3; ++ei_start) {
                        for (int ej_start = 0; ej_start < 3; ++ej_start) {
                            auto p11 = abi.world_V.row(fi[ei_start]);
                            auto p12 = abi.world_V.row(fi[(ei_start+1) % 3]);
                            auto p21 = abj.world_V.row(fj[ej_start]);
                            auto p22 = abj.world_V.row(fj[(ej_start+1) % 3]);
                            edge_edge_fn(i,abi.q,p11,p12,j,abj.q,p21,p22);
                        }
                    }
                }
            }
        }
    }

}


double contact_ip(Eigen::VectorXd q, std::vector<AffineBody> abs) {
    double ip = 0.0;
    iterate_primitives(q, abs,
        [&ip](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3) {
            double d = d2_vertex_face(q1,p,q2,p1,p2,p3);
            if (d >= dh) {
                return;
            }
            ip += ip_vertex_face(q1,p,q2,p1,p2,p3);
        },
        [&ip](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p) {
            double d = d2_vertex_face(q2,p,q1,p1,p2,p3);
            if (d >= dh) {
                return;
            }
            ip += ip_vertex_face(q2,p,q1,p1,p2,p3);
        },
        [&ip] (int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p11,
            Eigen::Vector3d p12,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p21,
            Eigen::Vector3d p22) {
            double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
            if (d >= dh) {
                return;
            }
            ip += ip_edge_edge(q1,p11,p12,q2,p21,p22);
        }
    );
    return ip;
}

Eigen::VectorXd contact_ip_gradient(Eigen::VectorXd q, std::vector<AffineBody> abs) {
    Eigen::VectorXd grad = Eigen::VectorXd::Zero(abs.size()*12);
    iterate_primitives(q, abs,
        [&grad](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3) {
            double d = d2_vertex_face(q1,p,q2,p1,p2,p3);
            if (d >= dh) {
                return;
            }
            auto grad_vf = gradient_vertex_face(q1,p,q2,p1,p2,p3);
            grad(Eigen::seqN(i*12,12)) += grad_vf(Eigen::seqN(0,12));
            grad(Eigen::seqN(j*12,12)) += grad_vf(Eigen::seqN(12,12));
        },
        [&grad](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p) {
            double d = d2_vertex_face(q2,p,q1,p1,p2,p3);
            if (d >= dh) {
                return;
            }
            auto grad_fv = gradient_vertex_face(q2,p,q1,p1,p2,p3);
            grad(Eigen::seqN(i*12,12)) += grad_fv(Eigen::seqN(12,12));
            grad(Eigen::seqN(j*12,12)) += grad_fv(Eigen::seqN(0,12));
        },
        [&grad] (int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p11,
            Eigen::Vector3d p12,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p21,
            Eigen::Vector3d p22) {
            double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
            if (d >= dh) {
                return;
            }
            auto grad_ee = gradient_edge_edge(q1,p11,p12,q2,p21,p22);
            grad(Eigen::seqN(i*12,12)) += grad_ee(Eigen::seqN(0,12));
            grad(Eigen::seqN(j*12,12)) += grad_ee(Eigen::seqN(12,12));
        }
    );
    return grad;
}

Eigen::MatrixXd contact_ip_hessian(Eigen::VectorXd q, std::vector<AffineBody> abs) {

    Eigen::MatrixXd hess = Eigen::MatrixXd::Zero(12*abs.size(),12*abs.size());
    iterate_primitives(q, abs,
        [&hess](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3) {
            double d = d2_vertex_face(q1,p,q2,p1,p2,p3);
            if (d >= dh) {
                return;
            }
            auto hess_fv = hessian_vertex_face(q1,p,q2,p1,p2,p3);
            hess.block<12,12>(i*12,i*12) += hess_fv.block<12,12>(0,0);
            hess.block<12,12>(j*12,j*12) += hess_fv.block<12,12>(12,12);
            hess.block<12,12>(i*12,j*12) += hess_fv.block<12,12>(0,12);
            hess.block<12,12>(j*12,i*12) += hess_fv.block<12,12>(12,0);
        },
        [&hess](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p) {
            double d = d2_vertex_face(q2,p,q1,p1,p2,p3);
            if (d >= dh) {
                return;
            }
            auto hess_fv = hessian_vertex_face(q2,p,q1,p1,p2,p3);
            hess.block<12,12>(i*12,i*12) += hess_fv.block<12,12>(12,12);
            hess.block<12,12>(j*12,j*12) += hess_fv.block<12,12>(0,0);
            hess.block<12,12>(i*12,j*12) += hess_fv.block<12,12>(12,0);
            hess.block<12,12>(j*12,i*12) += hess_fv.block<12,12>(0,12);
        },
        [&hess] (int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p11,
            Eigen::Vector3d p12,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p21,
            Eigen::Vector3d p22) {
            double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
            if (d >= dh) {
                return;
            }
            auto hess_ee = hessian_edge_edge(q1,p11,p12,q2,p21,p22);
            hess.block<12,12>(i*12,i*12) += hess_ee.block<12,12>(0,0);
            hess.block<12,12>(j*12,j*12) += hess_ee.block<12,12>(12,12);
            hess.block<12,12>(i*12,j*12) += hess_ee.block<12,12>(0,12);
            hess.block<12,12>(j*12,i*12) += hess_ee.block<12,12>(12,0);
        }
    );
    return hess;
}
