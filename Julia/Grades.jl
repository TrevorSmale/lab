### A Pluto.jl notebook ###
# v0.19.29

using Markdown
using InteractiveUtils

# ╔═╡ bab8fcff-e76e-4d5e-bc5c-4d318f47c286
import Pkg; Pkg.add("plots")

# ╔═╡ 50ac5366-2bd5-473d-b2fe-7a7ca3296cba
# 2. Import the Statistics module to access the `mean` function.
using Statistics

# ╔═╡ 7d09ecd5-9e0e-4a5b-ae47-948657c67eb6
using Plots

# ╔═╡ 086a4cb2-6fcf-11ee-176a-33850bc849a1
# Doing Stats in Julia

# ╔═╡ f2aced5f-220a-4de5-8af6-ef36521cb523
# 1. Define your set of grades as an array.
grades = [85, 92, 78, 90, 88]

# ╔═╡ cd4b1f99-50cd-4d40-bebf-ef18daf193e5
# 3. Calculate the mean using the `mean` function.
mean_grade = mean(grades)

# ╔═╡ 615f6b0e-2756-459c-8620-52435d8d6d46
# 4. Print or use the `mean_grade` as needed.
println("The mean grade is: $mean_grade")

# ╔═╡ 9a094a70-3734-4337-8273-1962d68e59f5
bar(grades, xlabel="Student", ylabel="Grades", title="Grades Distribution")

# ╔═╡ Cell order:
# ╠═bab8fcff-e76e-4d5e-bc5c-4d318f47c286
# ╠═086a4cb2-6fcf-11ee-176a-33850bc849a1
# ╠═f2aced5f-220a-4de5-8af6-ef36521cb523
# ╠═50ac5366-2bd5-473d-b2fe-7a7ca3296cba
# ╠═cd4b1f99-50cd-4d40-bebf-ef18daf193e5
# ╠═615f6b0e-2756-459c-8620-52435d8d6d46
# ╠═7d09ecd5-9e0e-4a5b-ae47-948657c67eb6
# ╠═9a094a70-3734-4337-8273-1962d68e59f5
