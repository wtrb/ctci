https://www.youtube.com/watch?v=KyUTuwz_b7Q
https://www.hackerearth.com/practice/data-structures/hash-tables/basics-of-hash-tables/tutorial/

## Hash table
- Linked list

## Hashing function/algorithms:
- Calculation applied to a key to transform it into an address
- For numeric keys, with n is the number of available addresses: `address = key MOD n`
- For alphanumeric keys, `address = sum ASCII codes MOD n`
- Folding method devides key into equal parts then add the parts together.
    + The telephone number: 0987654321, becomes 09 + 87 + 65 + 43 + 21 = 225
    + Depending on size of table, may then devide by some constant and take reminder

- Requirements:
    + Easy to compute: It should be easy to compute and must not become an algorithm in itself.
    + Uniform distribution: It should provide a uniform distribution across the hash table and should not result in clustering.
    + Less collisions: Collisions occur when pairs of elements are mapped to the same hash value. These should be avoided.

## Collisions
- Open addressing (separate chaining)
    + Linear probing
    + Plus 2 rehash
    + Quadratic probing (failed attempts)2
    + Double hashing
- Closed addressing


Homework:
- collision handling
- hashtables grow/shrink
- implement a simple hashtable
- practice questions
- hash functions