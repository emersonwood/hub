# codelingo/csharp ast lexicon

##  csharp facts


<details><summary>csharp.not_implemented</summary><p>

#### Example of finding every not_implemented and having a review bot comment on it:

```csharp
tenets:
  - name: find_all_not_implemented
    doc:  Example query to find all instances of not_implemented
    bots:
      codelingo/review
        comment: This is a not_implemented.
    query: |
      import codelingo/csharp

      @ review.comment
      csharp.not_implemented
```
</p></details>

