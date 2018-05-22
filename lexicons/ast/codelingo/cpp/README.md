# codelingo/cpp ast lexicon

##  cpp facts


<details><summary>cpp.not_implemented</summary><p>

#### Example of finding every not_implemented and having a review bot comment on it:

```cpp
tenets:
  - name: find_all_not_implemented
    doc:  Example query to find all instances of not_implemented
    bots:
      codelingo/review
        comment: This is a not_implemented.
    query: |
      import codelingo/cpp

      @ review.comment
      cpp.not_implemented
```
</p></details>

