#ifndef PLYGRD_AFFINEBODY_H
#define PLYGRD_AFFINEBODY_H

#include <chrono>
#include <iostream>
#include <vector>

#include <Eigen/Eigen>

struct AffineBody {
    const Eigen::Ref<const Eigen::MatrixX3d> world_V;
    const Eigen::Ref<const Eigen::MatrixX3i> F;

    Eigen::Vector<double,12> q;
    Eigen::Matrix<double,12,12> M;

    AffineBody(
        const Eigen::Ref<const Eigen::MatrixX3d> world_V,
        const Eigen::Ref<const Eigen::MatrixX3i> F,
        Eigen::Vector<double,12> q
    );

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
        const std::vector<AffineBody>& prev_states,
        std::chrono::milliseconds dt);

#endif  // PLYGRD_AFFINEBODY_H