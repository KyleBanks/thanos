# thanos

`thanos` is a simple little package that provides a `Snap` function, which accepts a slice and returns a subset containing half of the original elements, selected at random.

For example:

```
arr := []interface{}{1, 2, 3, 4, 5, 6}
out := thanos.Snap(arr)

println(out)
> [1, 3, 4]  // 3 random elements from the original array
```
