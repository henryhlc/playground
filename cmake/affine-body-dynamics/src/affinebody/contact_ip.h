#ifndef PLYGRD_CONTACTIP_H
#define PLYGRD_CONTACTIP_H

#include <Eigen/Eigen>

#include "affinebody.h"

double contact_ip(Eigen::VectorXd q, std::vector<AffineBody> abs);

Eigen::VectorXd contact_ip_gradient(Eigen::VectorXd q, std::vector<AffineBody> abs);

Eigen::MatrixXd contact_ip_hessian(Eigen::VectorXd q, std::vector<AffineBody> abs);

#endif  // PLYGRD_CONTACTIP_H