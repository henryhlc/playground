from sympy import *

q_1,q_2,q_3,q_4,q_5,q_6,q_7,q_8,q_9,q_10,q_11,q_12 = symbols(
    'q1_1 q1_2 q1_3 q1_4 q1_5 q1_6 q1_7 q1_8 q1_9 q1_10 q1_11 q1_12')


q1,q2,q3,q4,q5,q6,q7,q8,q9,q10,q11,q12 = symbols(
    'q1 q2 q3 q4 q5 q6 q7 q8 q9 q10 q11 q12'
)

q = Matrix([q1,q2,q3,q4,q5,q6,q7,q8,q9,q10,q11,q12])
p = Matrix([q1,q2,q3])
A = Matrix([[q4,q5,q6],[q7,q8,q9],[q10,q11,q12]])

def vector_code(var, expr):
    _,n = expr.shape
    init = f"    Eigen::VectorXd {var}{{{n}}};\n"
    fill = "\n".join([
        f"    {var}({i}) = {ccode(expr[i])};"
        for i in range(n)
    ])
    return init + fill

def matrix_code(var, expr):
    rs,cs = expr.shape
    init = f"    Eigen::MatrixXd {var}{{{rs},{cs}}};\n"
    fill = "\n".join([
        f"    {var}({r},{c}) = {ccode(expr[r,c])};"
        for r in range(rs)
        for c in range(cs)
    ])
    return init + fill

def q_extractor(in_var, out_prefix):
    return '\n'.join([
        f"    double {out_prefix}{i+1} = {in_var}[{i}];"
        for i in range(0,12)
    ]) +  '\n'

diff = A*A.T-eye(3)
ip = sum(diff.multiply_elementwise(diff))
ip_grad = Matrix([ip]).jacobian(q)
ip_hess = hessian(Matrix([ip]),q)

print("""
#include <Eigen/Eigen>
""")

print("""
double ip_orthogonal(Eigen::Vector<double,12> q) {{
{q_ex}
    return {expr};
}}
""".format(q_ex=q_extractor('q','q'),
           expr=ccode(ip)))

print("""
Eigen::Vector<double,12> grad_orthogonal(Eigen::Vector<double,12> q) {{
{q_ex}
{expr}
    return grad;
}}
""".format(q_ex=q_extractor('q','q'),
           expr=vector_code('grad',ip_grad)))

print("""
Eigen::Matrix<double,12,12> hess_orthogonal(Eigen::Vector<double,12> q) {{
{q_ex}
{expr}
    return hess;
}}
""".format(q_ex=q_extractor('q','q'),
           expr=matrix_code('hess',ip_hess)))