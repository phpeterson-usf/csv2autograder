# csv2autograder

This is a big hack to automatically upload (to canvas) code review data expressed in the CSV of a Google form. 

Assumptions
1. The students email addresses are given as name@dons.usfca.edu. Of course some students used their personal gmail accounts to upload the form
1. autograder upload expects to get github IDs rather than login IDs, so I
hacked autograder (`actions/upload.toml`to not do github-to-login ID mapping, 
since the login ID is already in the gmail address collected by the form
1. The current autograder barfs when test cases don't exist, even though
`grade upload` doesn't need test cases, so I faked up dummy test cases
1. In Canvas assignment, the assignment name is "project06-code-review".
Therefore, the fake test cases have to be named the same: `tests/project06-code-review/project06-code-review.toml`.
1. I uploaded using `grade upload -p project06-code-review`

If you're thinking it would have been easier to copy and paste the scores
by hand, that's my conclusion too!