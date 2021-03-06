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

2.1-4 考虑把两个n位二进制整数加起来的问题，这两个整数分别存储在两个n元数组A和B中。这两个整数的和应按二进制形式存储在一个(n+1)元数组C中。请给出该
     问题的形式化描述，并写出伪代码。

     答：


2.2-2 考虑排序存储在数组A中的n个数： 首先找出A中的最小元素并将其与A[1]中的元素进行交换。接着，找出A中的次最小元素并将其与A[2]中的元素进行交换。对
     A中前n-1个元素按该方式继续。该算法成为选择算法，写出其伪代码。该算法维持的循环不变式是什么？为什么它只需对前n-1个元素，而不是对所有n个元素
     运行？用theta记号给出选择排序的最好情况与最坏情况运行时间。

     答：代码见choose.go -> Choose
        循环不变式是前k项已经是最小的k项,并且前k项是按顺序排列的。
        只对前n-1个元素排序的原因是，如果前n-1个元素已经完成了选择排序，那么剩下的最后一个元素，一定是最大的元素，需要排在数组的最后。
        该算法无论是最好情况还是最坏情况，都需要在遍历整个数组的过程中，挨个对比“其后元素”中是否存在比当前元素更小的元素，因此最好情况和
        最坏情况都是O(n^2)的。

2.2-3 再次考虑现行查找问题（参见联系2.1-3）。假定要查找的元素等可能地为数组中的任意元素，平均需要检查输入序列的多少元素？最坏情况又如何？
      用theta记号给出现行查找的平均情况和最坏情况运行时间。证明你的答案。

      答: 不妨设数组中的元素个数为n，一般地，假设数组中的元素不重复。那么被查找的元素b等可能地成为数组中a的任意元素，那么b成为a[j]的概率
          为1/n，对于任意1<= j <= a.length。因此被查找次数的期望就是 1 * 1/n + 2 * 1/n + ... + n * 1/n = (n+1)/2，即平均需要
          检查(n+1)/2次。最坏情况是需要检查到最后一个元素，也就是n。查找的平均情况和最坏情况都是O(n)。


2.3-2 重写过程MERGE，使之不使用哨兵，而是一旦数组L或R的所有元素均被复制回A就立刻停止，然后把另一个数组的声誉部分复制回A。

      答: 见merge-sort.go -> merge


2.3-3 使用数学归纳法证明：当n刚好是2的幂时，以下递归式的解是T(n)=nlgn

      答：当k=1时有n=2时，显然成立。
         假设当k=l-1时，有T(2^{k-1})=2T(2^{k-2}) + 2^{k-1} = 2^{k-1} * (k-1)
         那么当k=l时，T(2^k) = 2T(2^{k-1}) + 2^k = 2*(2^{k-1} * (k-1)) + 2^k = 2^k * k,得证。

2.3-4 我们可以把插入排序表示为如下的一个递归过程。为了排序A[1..n]，我们递归地排序A[1..n-1]，然后把A[n]插入已排序的数组A[1..n-1]中，
     为插入排序的这个递归版本的最坏情况运行时间写一个递归式。

     答：代码见insertion-sort.go -> InsertionSortRecursion

2.3-5 回顾查找问题（参见联系2.1-3），注意到，如果序列A已排好序，就可以将该序列的中点与v进行比较。根据比较的结果，原序列中有一半就可以不用
      再做进一步的考虑了。二分查找算法重复这个过程，每次都将序列声誉部分的规模减半。为二分查找写出迭代递归的伪代码。证明：二分查找的最坏情况
      运行时间为O(lgn)。

      答：最坏的情况出现在被查找的元素小于等于序列的最小值或大于等于序列的最大值，很显然运行时间为O(lgn)。

2.3-6 注意到2.1节中的过程INSERTION-SORT的第5～7行的while循环采用一种线性查找来（反向）扫描已排好序的子数组A[1..j-1]。我们可以使用二分查找
      来把插入排序的最坏情况总运行时间改进到O（nlgn）吗？

      答：不一定。
          按照一般理解来说，二分查找最坏的情况的时间运行是O(lgn)，而进行n次二分查找，显然应该是O(nlgn)。二分查找可以使对比次数减少不少，由原来
          的O(n)到O(lgn)，然而对于插入排序而言，即使找到相应的下标，还要进行值的移动。当然对于链表而言，这个时间花费是廉价的，是可以将最坏的情况
          改进到O(nlgn)，但是对顺序表来说，使用二分查找并不能。


2-1 (在归并排序中对小数组采用插入排序) 虽然归并排序的最坏情况运行时间为O(nlgn)，而插入排序的最坏运行时间为O(n^2)，但是插入排序中的常量银子可能
    使它在n较小时，在许多机器上世纪运行得更快。因此，在归并排序中当自问题变得足够小时，采用插入排序来使递归的叶变粗是有意义的。考虑对归并排序的一
    种修改，其中使用插入排序来排序长度为k的n/k子表，然后使用标准的合并机制来合并这些子表，这里k是一个待定的值。
    a. 证明：插入排序最坏情况可以在O(nk)时间内排序每个长度为k的n/k个子表。
    b. 表明在最坏情况下如何在O(nlg(n/k))时间内合并这些子表。
    c. 假定修改后的算法的最坏情况运行时间为O(nk+nlg(n/k))，要使修改后的算法与标准的归并排序具有相同的运行时间，作为n的一个函数，借助O记号，
       k的最大值是什么。
    d. 在实践中，我们应该如何选择k？

    答：a. 对于单个长度为k的子表，最坏情况下时间复杂度为O(k^2)，那么n/k个子表最坏的时间复杂度就是O(k^2*n/k) = O(nk)。
       b. 对于n/k个子表两两合并，直到合并至只有一个，需要合并lgn次，而单词合并的时间复杂度为O(n)，所以总的时间复杂度是O(nlg(n/k))。
       c. 不知道。
       d.