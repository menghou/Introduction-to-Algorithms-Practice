2.1-2 重写INSERTION-SORT, 使之按非升序（而不是降序）排序。

    答:see insertion-sort.go >> InsertionSortIntReverse

2.1-3 考虑以下查找问题：
    输入：n个数的一个序列A=<a_1, a_2, ..., a_n>和一个值v.
    输出：下表i使得v=A[i]或者当v不在A中出现时，v为特殊值NIL。
    写出现行查找的伪代码，它扫描整个序列来查找v。使用一个循环不变式来证明你的算法是正确的。确保你的额循环不变式满足三条必要的性质。

    答：代码见 linear-search.go >> LinearSearch
    循环不变式->
    初始化：在循环开始之前，indexOutPut被初始化为NIL，而此时index = 0, a[:index]为空集合，所以并不存在a[i]==b,故此时indexOutPut==NIL
           是正确的。这表明第一次循环迭代之前循环不变式成立。
    保持：  不妨假设在当index = j - 1时，循环仍然为结束。即在前j次循环中，并没有触发循环体中的if语句，即a[a:j-1]中不存在下标i，使a[i]==b,
           此时，indexOutPut仍然保持indexOutPut==NIL。那么在index=j时，如果a[j]!=b，indexOutPut将仍然保持indexOutPut==NIL，也就是
           循环体的下一次迭代保持循环式不变；如果有a[j]==b，将触发循环终止。
    终止：  循环终止有两种情况：由if语句主动跳出循环，此时indexOutPut应该等于跳出循环时的索引值，也就是第一次有a[index]==b时的索引index，
           被赋值给了indexOutPut，保证了算法的正确性；在index循环至a.length-1时，依旧未触发if语句，这代表在遍历了整个数组之后，仍然没有
           找到满足条件的索引，而此时indexOutPut保持了初始化时的值NIL，故此时输出特殊值NIL，算法依旧正确。

2.14 考虑把两个n位二进制整数加起来的问题，这两个整数分别存储在两个n元数组A和B中。这两个整数的和应按二进制形式存储在一个(n+1)元数组C中。请给出该
     问题的形式化描述，并写出伪代码。

     答：