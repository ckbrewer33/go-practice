# **Data Structures in Go – Practice Curriculum**

## **Overview**

This curriculum is designed to rebuild fluency in Go syntax while reinforcing core data structures and algorithmic thinking. Each section includes:

- Concept refresher
- Implementation goals
- Practice exercises
- Key considerations (performance, tradeoffs)

---

## **1. Dynamic Array (ArrayList)**

### **Concept**

Go slices already behave like dynamic arrays, but implementing one manually helps reinforce how resizing and memory management work.

A dynamic array maintains:

- **Length**: number of elements in use
- **Capacity**: allocated space

When capacity is exceeded, a new larger array is allocated and elements are copied over (typically doubling capacity).

### **Implementation Goals**

- Append
- Insert
- Remove
- Get / Set
- Len / Cap

### **Key Ideas**

- Amortized O(1) append
- Copying cost during resize
- Bounds checking

---

## **2. Singly Linked List**

### **Concept**

A linked list is a chain of nodes where each node points to the next.

```
[Value | Next] -> [Value | Next] -> nil
```

Unlike arrays, elements are not stored contiguously in memory.

### **Pros / Cons**

**Pros**

- O(1) insertion at head
- No resizing cost

**Cons**

- O(n) access
- Extra memory for pointers

### **Implementation Goals**

- PushFront / PushBack
- PopFront
- Find
- Remove
- Len

### **Key Ideas**

- Pointer manipulation
- Edge cases (empty list, single node)

### **Stretch**

- Reverse in-place
- Cycle detection (Floyd’s algorithm)

---

## **3. Doubly Linked List**

### **Concept**

Each node points both forward and backward:

```
[Prev | Value | Next]
```

This allows efficient removal of arbitrary nodes.

### **Implementation Goals**

- PushFront / PushBack
- RemoveNode
- Forward and backward traversal

### **Key Ideas**

- Maintaining head and tail
- Updating both pointers correctly

---

## **4. Stack and Queue**

### **Stack (LIFO)**

Last In, First Out:

```
Push -> Pop
```

Use cases:

- Function calls
- Expression evaluation

### **Queue (FIFO)**

First In, First Out:

```
Enqueue -> Dequeue
```

Use cases:

- BFS
- Task scheduling

### **Implementation Goals**

Stack:

- Push
- Pop
- Peek

Queue:

- Enqueue
- Dequeue
- Peek

### **Key Ideas**

- Backed by slice or linked list
- O(1) operations (if implemented correctly)

---

## **5. Binary Search Tree (BST)**

### **Concept**

A tree where:

- Left subtree < node
- Right subtree > node

```
        8
       / \
      3   10
```

### **Operations**

- Insert
- Search
- Delete

### **Traversals**

- InOrder (sorted output)
- PreOrder
- PostOrder

### **Key Ideas**

- Recursive structure
- Deletion cases:
    - Leaf
    - One child
    - Two children (replace with successor)

### **Complexity**

- Average: O(log n)
- Worst (unbalanced): O(n)

---

## **6. Heap (Priority Queue)**

### **Concept**

A heap is a complete binary tree stored in an array.

For a **min-heap**:

- Parent <= children

### **Representation**

For index `i`:

- Left child: `2i + 1`
- Right child: `2i + 2`
- Parent: `(i - 1) / 2`

### **Operations**

- Push (heapify up)
- Pop (heapify down)

### **Key Ideas**

- Maintaining heap invariant
- Efficient top element retrieval (O(1))

### **Complexity**

- Insert: O(log n)
- Remove: O(log n)

---

## **7. Hash Map**

### **Concept**

Maps keys to values using a hash function.

```
hash(key) -> index -> bucket
```

### **Collision Handling**

- Separate chaining (linked lists per bucket)

### **Implementation Goals**

- Put
- Get
- Delete
- Resize

### **Key Ideas**

- Load factor
- Rehashing
- Uniform distribution of keys

### **Complexity**

- Average: O(1)
- Worst: O(n)

---

## **8. Sorting Algorithms**

### **Bubble Sort**

- Repeatedly swap adjacent elements
- O(n²)

### **Insertion Sort**

- Build sorted list incrementally
- Good for small or nearly sorted data

### **Merge Sort**

- Divide and conquer
- Stable
- O(n log n)

### **Quick Sort**

- Partition around pivot
- Fast in practice
- Worst-case O(n²)

### **Heap Sort**

- Use heap to sort
- O(n log n)

### **Key Ideas**

- Stability
- In-place vs extra memory
- Tradeoffs depending on data

---

## **9. Graphs**

### **Concept**

A graph consists of vertices and edges.

Representation:

- Adjacency list (map → slice)

### **Traversals**

**BFS (Breadth-First Search)**

- Uses queue
- Explores level by level

**DFS (Depth-First Search)**

- Uses recursion or stack

### **Key Ideas**

- Visited tracking
- Directed vs undirected

### **Applications**

- Shortest path (unweighted)
- Connectivity

---

## **10. LRU Cache (Capstone)**

### **Concept**

Least Recently Used cache evicts the oldest unused item.

### **Design**

Combine:

- Hash map → O(1) lookup
- Doubly linked list → O(1) ordering

### **Operations**

- Get
- Put

### **Behavior**

- Access moves item to front
- Insert at front
- Evict from tail

### **Key Insight**

This is a classic example of **composing data structures** to meet strict performance constraints.

---

## **Recommended Workflow**

For each structure:

1. Write interface first
2. Implement incrementally
3. Add table-driven tests
4. Add benchmarks (where relevant)
5. Document Big-O complexity

---

## **Final Note**

Focus on:

- Writing without autocomplete/AI initially
- Explaining the structure out loud (teaching mindset)
- Understanding tradeoffs, not just implementation

This combination will rebuild both syntax fluency and deep intuition.