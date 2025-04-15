#set page(margin: 0.8cm)
#set text(font: "Charter", size: 10pt)
#set page(paper: "a4")
#set page(footer: context [
  Name Resume
  #h(0.5fr)
  Page
  #counter(page).display(
    "1 of 1",
    both: true,
  )
])

#let name = ("Name")
#let repo = link("https://github.com/yourrepo", "github.com/yourrepo")
#let email = ("example@email.com")
#let blog = link("https://website.io", "https://website.io")
#let summary = ("2 Sentence statement.")

#v(-10pt)
#align(center)[
= #name
#text(size: 10pt)[
#repo\
#email\
]]

#align(center)[
#line(length: 110%)
]

== *Summary*

#summary

== *Certifications*

*Cert1*, *Cert2*, *Cert3*

== *Skills*
*Skill1*, *Skill2*, *Sill3* \

== *Experience*

*Job Title* - Location

#grid(
  columns: (auto, 1fr),
  gutter: 1em,
 [Company],
  align(right)[2024 - Current]
)

// STAR Method description ----------------------------

- STAR
- STAR
- STAR
- STAR

//Begin Job Position 2

*Job Title* - Location

#grid(
  columns: (auto, 1fr),
  gutter: 1em,
 [Company],
  align(right)[2024 - Current]
)

// STAR Method description ----------------------------

- STAR
- STAR
- STAR
- STAR

//End Job Position 2

//Begin Job Position 3

*Job Title* - Location

#grid(
  columns: (auto, 1fr),
  gutter: 1em,
 [Company],
  align(right)[2018 - 2022]
)

- STAR
- STAR
- STAR
- STAR

//End Job Position 3
//End Experience

== *Education*

- Masters of Mischief - University of Homeschool
- Bachelor of Tickling - Laughington Estate University

#pagebreak()

//Projects Beginning
= *Projects*

#grid(
  columns: (auto, 1fr),
  gutter: 1em,
 [*Project 1 ⚓️*],
 [#text(rgb("#5c2815"))[#link("https://github.com/LinkToRepo")[https://github.com/LinkToRepo]]],
)

- STAR
- STAR
- STAR
- STAR

#grid(
  columns: (auto, 1fr),
  gutter: 1em,
 [*Project 2 ⚓️*],
 [#text(rgb("#5c2815"))[#link("https://github.com/LinkToRepo")[https://github.com/LinkToRepo]]],
)

- STAR
- STAR
- STAR
- STAR

#grid(
  columns: (auto, 1fr),
  gutter: 1em,
 [*Project 3 ⚓️*],
 [#text(rgb("#5c2815"))[#link("https://github.com/LinkToRepo")[https://github.com/LinkToRepo]]],
)

- STAR
- STAR
- STAR
- STAR

