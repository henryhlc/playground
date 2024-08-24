# Affine body dynamics

- Code generation: sympy
- Linked libraries: SDL3, Eigen3, libigl, polyscope
- Data sources: alecjacobson/computer-graphics-meshes

References
- (2022) Affine Body Dynamics: Fast, Stable & Intersection-free Simulation of Stiff Materials

Known issues
- dependencies that uses `Eigen::all` needs to be updated to `Eigen::indexing::all` or build will fail.
