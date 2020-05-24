# Recursion

### To solve a problem using recursion, then we need to make sure that:
- The problem can broken down into smaller problems of same type.
- Problem has some base case(s).
- Base case is reached before the stack size limit exceeds.

* Base case:
Any recursive method must have a `terminating condition`. Terminating condition is one for which the answer is `already known` and we just need to return that. For example for the factorial problem, we know that `factorial(0) = 1`


## Backtracking
Algorithm, tries solving a subproblem, if that does not result in the solution, it undo whatever changes were made and solve the next subproblem. If the solution does not exists, then it returns false.



https://www.hackerearth.com/practice/basic-programming/recursion/recursion-and-backtracking/tutorial/