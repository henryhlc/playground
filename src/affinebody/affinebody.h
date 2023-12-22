#ifndef PLYGRD_AFFINEBODY_H
#define PLYGRD_AFFINEBODY_H

#include <chrono>
#include <iostream>
#include <vector>

#include <Eigen/Eigen>

struct AffineBody {
    Eigen::MatrixX3d world_V;
    Eigen::MatrixX3i F;

    Eigen::Vector<double,12> q;
    Eigen::Matrix<double,12,12> M;

    AffineBody(
        const Eigen::MatrixX3d world_V,
        const Eigen::MatrixX3i F,
        const Eigen::Vector<double,12> q
    );

    AffineBody(const AffineBody& other): AffineBody(other.world_V, other.F, other.q) {};
    AffineBody& operator=(const AffineBody& other) = default;
        // if (this != &other) {
            // this->q = other.q;
            // this->M = other.M;
            // this->world_V = other.world_V;
            // this->F = other.F;
        // }
        // return *this;
    // }

    Eigen::Vector3d p() const {
        return q(Eigen::seq(0,2));
    }

    Eigen::Matrix3d A() const {
        Eigen::Matrix3d res;
        res.row(0) = q(Eigen::seq(3, 5));
        res.row(1) = q(Eigen::seq(6, 8));
        res.row(2) = q(Eigen::seq(9, 11));
        return res;
    }

    Eigen::MatrixX3d V() const {
        return (world_V * A().transpose()).rowwise() + p().transpose();
    }
};

std::vector<AffineBody> affine_body_dynamics(
        const std::vector<AffineBody>& curr_states,
        const std::vector<AffineBody>& prev_states,
        std::chrono::milliseconds dt);

#endif  // PLYGRD_AFFINEBODY_H