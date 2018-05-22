# codelingo/python27 ast lexicon

##  python27 facts


<details><summary>python27.not_implemented</summary><p>

#### Example of finding every not_implemented and having a review bot comment on it:

```python27
tenets:
  - name: find_all_not_implemented
    doc:  Example query to find all instances of not_implemented
    bots:
      codelingo/review
        comment: This is a not_implemented.
    query: |
      import codelingo/python27

      @ review.comment
      python27.not_implemented
```
</p></details>

