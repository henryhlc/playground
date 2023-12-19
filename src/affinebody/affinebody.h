#ifndef PLYGRD_AFFINEBODY_H
#define PLYGRD_AFFINEBODY_H

#include <chrono>
#include <iostream>
#include <vector>

#include <Eigen/Eigen>

#include "./massmatrix_integral.h"

Eigen::Matrix<double, 12, 12> mass_matrix(
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

struct AffineBody {
    Eigen::Vector<double,12> q;

    const Eigen::Ref<const Eigen::MatrixX3d> world_V;
    const Eigen::Ref<const Eigen::MatrixX3i> F;

    AffineBody(
        const Eigen::Ref<const Eigen::MatrixX3d> world_V,
        const Eigen::Ref<const Eigen::MatrixX3i> F,
        Eigen::Vector<double,12> q
    ): world_V(world_V), F(F), q(q) {}
    
    Eigen::Vector3d p() {
        return q(Eigen::seq(0,2));
    }

    Eigen::Matrix3d A() {
        Eigen::Matrix3d res;
        res.row(0) = q(Eigen::seq(3, 5));
        res.row(1) = q(Eigen::seq(6, 8));
        res.row(2) = q(Eigen::seq(9, 11));
        return res;
    }

    Eigen::MatrixX3d V() {
        return (world_V * A().transpose()).rowwise() + p().transpose();
    }
};

std::vector<AffineBody> affine_body_dynamics(
        const std::vector<AffineBody>& curr_states,
        std::chrono::milliseconds dt) {
    std::vector<AffineBody> next_states {}; 
    for (auto& af : curr_states) {
        auto q_next = af.q;
        q_next(Eigen::seq(0,2)) += Eigen::Vector3d {0.01, 0.01, 0.01};
        next_states.emplace_back(af.world_V, af.F, q_next);
    }
    return next_states;
}

#endif  // PLYGRD_AFFINEBODY_H