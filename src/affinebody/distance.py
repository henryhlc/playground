from sympy import *

# Keep in sync with the value in distance.h
dh = 0.03

q1_1,q1_2,q1_3,q1_4,q1_5,q1_6,q1_7,q1_8,q1_9,q1_10,q1_11,q1_12 = symbols(
    'q1_1 q1_2 q1_3 q1_4 q1_5 q1_6 q1_7 q1_8 q1_9 q1_10 q1_11 q1_12')

q2_1,q2_2,q2_3,q2_4,q2_5,q2_6,q2_7,q2_8,q2_9,q2_10,q2_11,q2_12 = symbols(
    'q2_1 q2_2 q2_3 q2_4 q2_5 q2_6 q2_7 q2_8 q2_9 q2_10 q2_11 q2_12')

q1 = Matrix([q1_1,q1_2,q1_3,q1_4,q1_5,q1_6,q1_7,q1_8,q1_9,q1_10,q1_11,q1_12])
p1 = Matrix([q1_1,q1_2,q1_3])
A1 = Matrix([[q1_4,q1_5,q1_6],[q1_7,q1_8,q1_9],[q1_10,q1_11,q1_12]])
q2 = Matrix([q2_1,q2_2,q2_3,q2_4,q2_5,q2_6,q2_7,q2_8,q2_9,q2_10,q2_11,q2_12])
p2 = Matrix([q2_1,q2_2,q2_3])
A2 = Matrix([[q2_4,q2_5,q2_6],[q2_7,q2_8,q2_9],[q2_10,q2_11,q2_12]])
q = Matrix([q1,q2])

# Point-vertex pair

x,y,z = symbols('x y z')
x1,y1,z1,x2,y2,z2,x3,y3,z3 = symbols('x1 y1 z1 x2 y2 z2 x3 y3 z3')

P = p1 + A1 * Matrix([x,y,z])
P1 = p2 + A2 * Matrix([x1,y1,z1])
P2 = p2 + A2 * Matrix([x2,y2,z2])
P3 = p2 + A2 * Matrix([x3,y3,z3])

# Edge-edge pair
x11,y11,z11,x12,y12,z12 = symbols('x11 y11 z11 x12 y12 z12')
x21,y21,z21,x22,y22,z22 = symbols('x21 y21 z21 x22 y22 z22')
P11 = p1 + A1 * Matrix([x11,y11,z11])
P12 = p1 + A1 * Matrix([x12,y12,z12])
P21 = p2 + A2 * Matrix([x21,y21,z21])
P22 = p2 + A2 * Matrix([x22,y22,z22])

d = symbols('d')
barrier = -(dh-d)**2*log(d/dh)

def grad(v, q):
    return Matrix([v]).jacobian(q)

def hess(v,q):
    return hessian(Matrix([v]), q)


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
        f"    double {out_prefix}_{i+1} = {in_var}[{i}];"
        for i in range(0,12)
    ]) +  '\n'

def v3_extractor(in_var, out_suffix):
    return """    double x{out_suffix} = {in_var}(0);
    double y{out_suffix} = {in_var}(1);
    double z{out_suffix} = {in_var}(2);
""".format(
    out_suffix=out_suffix,
    in_var=in_var
)

def d2_grad_hess_fns(fns, expr):
    return '\n'.join([
    fns('double', '', f"    return {ccode(expr)};"),
    fns('Eigen::VectorXd', '_grad', """
{init_and_fill}
    return grad;""".format(
        init_and_fill=vector_code('grad', grad(expr,q))
    )),
    fns('Eigen::MatrixXd', '_hess', """
{init_and_fill}
    return hess;""".format(
        init_and_fill=matrix_code('hess', hess(expr,q))
    ))
    ])

# point-point distance squared
def d2_pp(p1,p2):
    s12 = p2 - p1
    return (s12.T * s12)[0,0]

d2_pp_expr = d2_pp(P,P1)

def d2_pp_fns(type, suffix, content):
    return """
{type} d2_pp{suffix}(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d a,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d b
) {{
{q1_ex}
{q2_ex}
{x_ex}
{x1_ex}
{content}
}}
""".format(
    type=type,
    suffix=suffix,
    q1_ex=q_extractor('q1', 'q1'),
    q2_ex=q_extractor('q2', 'q2'),
    x_ex=v3_extractor('a', ''),
    x1_ex=v3_extractor('b', '1'),
    content=content,
)


# point-edge distance
def d2_pe(p,p1,p2):
    s12 = p2 - p1
    cp = (p1-p).cross(p2-p)
    return (cp.T*cp)[0,0] / (s12.T*s12)[0,0]

d2_pe_expr = d2_pe(P,P1,P2)

def d2_pe_fns(type, suffix, content):
    return """
{type} d2_pe{suffix}(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2) {{
{q1_ex}
{q2_ex}
{p_ex}
{p1_ex}
{p2_ex}
{content}
}}
""".format(
    type=type,
    suffix=suffix,
    q1_ex=q_extractor('q1', 'q1'),
    q2_ex=q_extractor('q2', 'q2'),
    p_ex=v3_extractor('p', ''),
    p1_ex=v3_extractor('p1', '1'),
    p2_ex=v3_extractor('p2', '2'),
    content=content
)

# point-triangle distance
def d2_pt(p,p1,p2,p3):
    cp = (p2-p1).cross(p3-p1)
    return (p-p1).dot(cp)**2 / (cp.T*cp)[0,0]
d2_pt_expr = d2_pt(P,P1,P2,P3)

def d2_pt_fns(type, suffix, content):
    return """
{type} d2_pt{suffix}(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p1,
    Eigen::Vector3d p2,
    Eigen::Vector3d p3) {{
{q1_ex}
{q2_ex}
{p_ex}
{p1_ex}
{p2_ex}
{p3_ex}
{content}
}}
""".format(
    type=type,
    suffix=suffix,
    q1_ex=q_extractor('q1', 'q1'),
    q2_ex=q_extractor('q2', 'q2'),
    p_ex=v3_extractor('p', ''),
    p1_ex=v3_extractor('p1', '1'),
    p2_ex=v3_extractor('p2', '2'),
    p3_ex=v3_extractor('p3', '3'),
    content=content
)



# edge-edge distance
def d2_ee(p11,p12,p21,p22):
    cp = (p12-p11).cross(p22-p21)
    return (p11-p21).dot(cp)**2 / (cp.T*cp)[0,0]

d2_ee_expr = d2_ee(P11,P12,P21,P22)

def d2_ee_fns(type, suffix, content):
    return """
{type} d2_ee{suffix}(
    Eigen::Vector<double,12> q1,
    Eigen::Vector3d p11,
    Eigen::Vector3d p12,
    Eigen::Vector<double,12> q2,
    Eigen::Vector3d p21,
    Eigen::Vector3d p22) {{
{q1_ex}
{q2_ex}
{p11_ex}
{p12_ex}
{p21_ex}
{p22_ex}
{content}
}}
""".format(
    type=type,
    suffix=suffix,
    q1_ex=q_extractor('q1', 'q1'),
    q2_ex=q_extractor('q2', 'q2'),
    p11_ex=v3_extractor('p11', '11'),
    p12_ex=v3_extractor('p12', '12'),
    p21_ex=v3_extractor('p21', '21'),
    p22_ex=v3_extractor('p22', '22'),
    content=content
)

print("""
// This is generated by distance.py.
#include <Eigen/Eigen>
#include <cmath>""")

print("""
double barrier(double d) {{
    return {b_expr};
}}

double d_barrier(double d) {{
    return {db_expr}
}}

double dd_barrier(double d) {{
    return {ddb_expr}
}}
""".format(
    b_expr=ccode(barrier),
    db_expr=ccode(diff(barrier,d)),
    ddb_expr=ccode(diff(diff(barrier,d),d))
))

print(d2_grad_hess_fns(d2_pt_fns, d2_pt_expr))
print(d2_grad_hess_fns(d2_ee_fns, d2_ee_expr))
print(d2_grad_hess_fns(d2_pe_fns, d2_pe_expr))
print(d2_grad_hess_fns(d2_pp_fns, d2_pp_expr))