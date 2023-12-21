#include "contact_ip.h"

#include <functional>

#include <ipc/distance/distance_type.hpp>

#include "./distance.h"

double d2_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    switch (ipc::point_triangle_distance_type(p,p1,p2,p3)) {
        case ipc::PointTriangleDistanceType::P_T0:
            return d2_pp(q1,p,q2,p1);
        case ipc::PointTriangleDistanceType::P_T1:
            return d2_pp(q1,p,q2,p2);
        case ipc::PointTriangleDistanceType::P_T2:
            return d2_pp(q1,p,q2,p3);
        case ipc::PointTriangleDistanceType::P_E0:
            return d2_pe(q1,p,q2,p1,p2);
        case ipc::PointTriangleDistanceType::P_E1:
            return d2_pe(q1,p,q2,p2,p3);
        case ipc::PointTriangleDistanceType::P_E2:
            return d2_pe(q1,p,q2,p3,p1);
        case ipc::PointTriangleDistanceType::P_T:
            return d2_pt(q1,p,q2,p1,p2,p3);
    }
    return 0.0;
}

double ip_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    double d {d2_vertex_face(q1,p,q2,p1,p2,p3)}
    if (d >= dh) {
        return 0;
    }
    return barrier(d);
}

Eigen::Vector<double,24> gradient_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    double d {d2_vertex_face(q1,p,q2,p1,p2,p3)}
    if (d >= dh) {
        return Eigen::Vector<double,24>::Zero();
    }
    Eigen::Vector<double,24> d_grad
    switch (ipc::point_triangle_distance_type(p,p1,p2,p3)) {
        case ipc::PointTriangleDistanceType::P_T0:
            d_grad = d2_pp_grad(q1,p,q2,p1);
        case ipc::PointTriangleDistanceType::P_T1:
            d_grad = d2_pp_grad(q1,p,q2,p2);
        case ipc::PointTriangleDistanceType::P_T2:
            d_grad = d2_pp_grad(q1,p,q2,p3);
        case ipc::PointTriangleDistanceType::P_E0:
            d_grad = d2_pe_grad(q1,p,q2,p1,p2);
        case ipc::PointTriangleDistanceType::P_E1:
            d_grad = d2_pe_grad(q1,p,q2,p2,p3);
        case ipc::PointTriangleDistanceType::P_E2:
            d_grad = d2_pe_grad(q1,p,q2,p3,p1);
        case ipc::PointTriangleDistanceType::P_T:
            d_grad = d2_pt_grad(q1,p,q2,p1,p2,p3);
    }
    return d_barrier(d) * d_grad;
}

Eigen::Matrix<double,24,24> hessian_vertex_face(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {
    double d {d2_vertex_face(q1,p,q2,p1,p2,p3)}
    if (d >= dh) {
        return Eigen::Matrix<double,24,24>::Zero();
    }
    Eigen::Matrix<double,24,24> d_hess;
    switch (ipc::point_triangle_distance_type(p,p1,p2,p3)) {
        case ipc::PointTriangleDistanceType::P_T0:
            d_hess = d2_pp_hess(q1,p,q2,p1);
        case ipc::PointTriangleDistanceType::P_T1:
            d_hess = d2_pp_hess(q1,p,q2,p2);
        case ipc::PointTriangleDistanceType::P_T2:
            d_hess = d2_pp_hess(q1,p,q2,p3);
        case ipc::PointTriangleDistanceType::P_E0:
            d_hess = d2_pe_hess(q1,p,q2,p1,p2);
        case ipc::PointTriangleDistanceType::P_E1:
            d_hess = d2_pe_hess(q1,p,q2,p2,p3);
        case ipc::PointTriangleDistanceType::P_E2:
            d_hess = d2_pe_hess(q1,p,q2,p3,p1);
        case ipc::PointTriangleDistanceType::P_T:
            d_hess = d2_pt_hess(q1,p,q2,p1,p2,p3);
    }
    Eigen::Vector<double,12> d_grad = gradient_vertex_face(q1,p,q2,p1,p2,p3);
    return d_barrier(d) * d_hess + dd_barrier(d) * d_grad * d_grad.transpose();
}

double d2_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {
    switch (ipc::edge_edge_distance_type(p11,p12,p21,p22)) {
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
    double d = d2_edge_edge(q1,p11,p12,q2,p21,p22) ;
    if (d >= dh) {
        return 0.0;
    }
    return barrier(d);
}

Eigen::Vector<double,24> gradient_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {
    double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
    if (d >= dh) {
        return Eigen::Vector<double,24>::Zero();
    }
    Eigen::Vector<double,12> d_grad;
    switch (ipc::edge_edge_distance_type(p11,p12,p21,p22)) {
        case ipc::EdgeEdgeDistanceType::EA0_EB0:
            d_grad = d2_pp_grad(q1,p11,q2,p21);
        case ipc::EdgeEdgeDistanceType::EA0_EB1:
            d_grad = d2_pp_grad(q1,p11,q2,p22);
        case ipc::EdgeEdgeDistanceType::EA1_EB0:
            d_grad = d2_pp_grad(q1,p12,q2,p21);
        case ipc::EdgeEdgeDistanceType::EA1_EB1:
            d_grad = d2_pp_grad(q1,p12,q2,p22);
        case ipc::EdgeEdgeDistanceType::EA_EB0: {
            Eigen::Vector<double,24> grad;
            auto pe_grad = d2_pe_grad(q2,p21,q1,p11,p12);
            grad(Eigen::seqN(0,12)) = pe_grad(Eigen::seqN(12,12));
            grad(Eigen::seqN(12,12)) = pe_grad(Eigen::seqN(0,12));
            d_grad = grad;
        }
        case ipc::EdgeEdgeDistanceType::EA_EB1: {
            Eigen::Vector<double,24> grad;
            auto pe_grad = d2_pe_grad(q2,p22,q1,p11,p12);
            grad(Eigen::seqN(0,12)) = pe_grad(Eigen::seqN(12,12));
            grad(Eigen::seqN(12,12)) = pe_grad(Eigen::seqN(0,12));
            d_grad = grad;
        }
        case ipc::EdgeEdgeDistanceType::EA0_EB:
            d_grad = d2_pe_grad(q1,p11,q2,p21,p22);
        case ipc::EdgeEdgeDistanceType::EA1_EB:
            d_grad = d2_pe_grad(q1,p12,q2,p21,p22);
        case ipc::EdgeEdgeDistanceType::EA_EB:
            d_grad = d2_ee_grad(q1,p11,p12,q2,p21,p22);
    }
    return d_barrier(d) * d_grad;
}

Eigen::Matrix<double,24,24> hessian_edge_edge(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {
    double d = d2_edge_edge(q1,p11,p12,q2,p21,p22);
    if (d >= dh) {
        return Eigen::Matrix<double,24,24>::Zero();
    }
    Eigen::Matrix<double,24,24> d_hess;
    switch (ipc::edge_edge_distance_type(p11,p12,p21,p22)) {
        case ipc::EdgeEdgeDistanceType::EA0_EB0:
            d_hess = d2_pp_hess(q1,p11,q2,p21);
        case ipc::EdgeEdgeDistanceType::EA0_EB1:
            d_hess = d2_pp_hess(q1,p11,q2,p22);
        case ipc::EdgeEdgeDistanceType::EA1_EB0:
            d_hess = d2_pp_hess(q1,p12,q2,p21);
        case ipc::EdgeEdgeDistanceType::EA1_EB1:
            d_hess = d2_pp_hess(q1,p12,q2,p22);
        case ipc::EdgeEdgeDistanceType::EA_EB0: {
            Eigen::Matrix<double,24,24> hess;
            auto pe_hess = d2_pe_hess(q2,p21,q1,p11,p12);
            hess.block<12,12>(0,0) = pe_hess.block<12,12>(12,12);
            hess.block<12,12>(12,0) = pe_hess.block<12,12>(0,12);
            hess.block<12,12>(0,12) = pe_hess.block<12,12>(12,0);
            hess.block<12,12>(12,12) = pe_hess.block<12,12>(0,0);
            d_hess = hess;
        }
        case ipc::EdgeEdgeDistanceType::EA_EB1: {
            Eigen::Matrix<double,24,24> hess;
            auto pe_hess = d2_pe_hess(q2,p22,q1,p11,p12);
            hess.block<12,12>(0,0) = pe_hess.block<12,12>(12,12);
            hess.block<12,12>(12,0) = pe_hess.block<12,12>(0,12);
            hess.block<12,12>(0,12) = pe_hess.block<12,12>(12,0);
            hess.block<12,12>(12,12) = pe_hess.block<12,12>(0,0);
            d_hess = hess;
        }
        case ipc::EdgeEdgeDistanceType::EA0_EB:
            d_hess = d2_pe_hess(q1,p11,q2,p21,p22);
        case ipc::EdgeEdgeDistanceType::EA1_EB:
            d_hess = d2_pe_hess(q1,p12,q2,p21,p22);
        case ipc::EdgeEdgeDistanceType::EA_EB:
            d_hess = d2_ee_hess(q1,p11,p12,q2,p21,p22);
    }
    Eigen::Vector<double,12> d_grad = gradient_edge_edge(q1,p11,p12,q2,p21,p22);
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
            auto Vi = abi.V();
            auto Vj = abj.V();

            // all pairs of faces
            for (auto& fi : abi.F.rowwise()) {
                for (auto& fj : abj.F.rowwise()) {
                    auto vi1 = Vi.row(fi[0]);
                    auto vi2 = Vi.row(fi[1]);
                    auto vi3 = Vi.row(fi[2]);
                    auto vj1 = Vj.row(fj[0]);
                    auto vj2 = Vj.row(fj[1]);
                    auto vj3 = Vj.row(fj[2]);

                    // all face-vertex
                    for (auto vj_index : fj) {
                        // fi and vj
                        Eigen::Vector3d vj = Vj.row(vj_index);
                        face_vertex_fn(i,abi.q,vi1,vi2,vi3,j,abj.q,vj);
                    }
                    // all vertex-face
                    for (auto vi_index : fi) {
                        // vi and fj
                        Eigen::Vector3d vi = Vi.row(vi_index);
                        vertex_face_fn(i,abi.q,vi,j,abj.q,vj1,vj2,vj3);

                    }
                    // all edge-edge
                    for (auto& vi_index : fi) {
                        for (auto& vj_index : fj) {
                            // vi-neighbor and vj-neighbor
                            auto p11 = Vi.row(vi_index);
                            auto p12 = Vi.row((vi_index + 1) % 3);
                            auto p21 = Vj.row(vj_index);
                            auto p22 = Vj.row((vj_index + 1) % 3);
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
            ip += ip_edge_edge(q1,p11,p12,q2,p21,p22);
        }
    );
    return ip;
}

Eigen::VectorXd contact_ip_gradient(Eigen::VectorXd q, std::vector<AffineBody> abs) {
    Eigen::VectorXd grad {abs.size()*12};
    iterate_primitives(q, abs,
        [&grad](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3) {
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
            auto grad_ee = gradient_edge_edge(q1,p11,p12,q2,p21,p22);
            grad(Eigen::seqN(i*12,12)) += grad_ee(Eigen::seqN(0,12));
            grad(Eigen::seqN(j*12,12)) += grad_ee(Eigen::seqN(12,12));
        }
    );
    return grad;
}

Eigen::MatrixXd contact_ip_hessian(Eigen::VectorXd q, std::vector<AffineBody> abs) {

    Eigen::MatrixXd hess{12*abs.size(),12*abs.size()};
    iterate_primitives(q, abs,
        [&hess](int i,
            Eigen::Vector<double,12> q1,
            Eigen::Vector3d p,
            int j,
            Eigen::Vector<double,12> q2,
            Eigen::Vector3d p1,
            Eigen::Vector3d p2,
            Eigen::Vector3d p3) {
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
            auto hess_ee = hessian_edge_edge(q1,p11,p12,q2,p21,p22);
            hess.block<12,12>(i*12,i*12) += hess_ee.block<12,12>(0,0);
            hess.block<12,12>(j*12,j*12) += hess_ee.block<12,12>(12,12);
            hess.block<12,12>(i*12,j*12) += hess_ee.block<12,12>(0,12);
            hess.block<12,12>(j*12,i*12) += hess_ee.block<12,12>(12,0);
        }
    );
    return hess;
}
