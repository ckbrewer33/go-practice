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

### **Method Notes**

- **Append(value)**: Add a value to the end of the array. If length equals capacity, allocate a larger backing slice, copy the existing values, then store the new value.
- **Insert(index, value)**: Add a value at a specific position. Check bounds, grow if needed, shift values from that index to the right, then place the new value.
- **Remove(index)**: Remove the value at a specific position. Check bounds, save the removed value if returning it, shift later values left, clear the unused final slot, then decrement length.
- **Get(index)**: Return the value at a specific position after checking that the index is within the current length.
- **Set(index, value)**: Replace the value at a specific position after checking that the index is within the current length.
- **Len()**: Return the number of values currently stored.
- **Cap()**: Return the amount of allocated space available before another resize is required.

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

### **Method Notes**

- **PushFront(value)**: Add a new node before the current head. Update the head pointer and handle the empty-list case.
- **PushBack(value)**: Add a new node after the current tail or walk to the end if no tail pointer is stored. Update the tail pointer if your list tracks one.
- **PopFront()**: Remove and return the current head value. Move head to the next node and handle empty-list and single-node cases.
- **Find(value)**: Walk node by node from the head until the value is found or the list ends.
- **Remove(value)**: Remove the first node containing the value. Track the previous node so its `Next` pointer can skip the removed node.
- **Len()**: Return the number of nodes currently in the list, either from a stored counter or by walking the list.

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

### **Method Notes**

- **PushFront(value)**: Add a node before the current head. Update the old head's `Prev`, the new node's `Next`, and the list's head pointer.
- **PushBack(value)**: Add a node after the current tail. Update the old tail's `Next`, the new node's `Prev`, and the list's tail pointer.
- **RemoveNode(node)**: Unlink a known node by connecting its previous and next neighbors to each other. Update head or tail if the removed node was at either end.
- **TraverseForward()**: Start at head and follow `Next` pointers until the end.
- **TraverseBackward()**: Start at tail and follow `Prev` pointers until the beginning.

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

### **Method Notes**

- **Stack Push(value)**: Add a value to the top of the stack.
- **Stack Pop()**: Remove and return the most recently pushed value. Return an error or second boolean when the stack is empty.
- **Stack Peek()**: Return the top value without removing it.
- **Queue Enqueue(value)**: Add a value to the back of the queue.
- **Queue Dequeue()**: Remove and return the oldest value from the front of the queue. Avoid repeatedly shifting a slice if you want true O(1) behavior.
- **Queue Peek()**: Return the front value without removing it.

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

### **Method Notes**

- **Insert(value)**: Compare the value against each node and move left or right until an empty child position is found.
- **Search(value)**: Compare the target value against each node and follow the left or right subtree until found or nil.
- **Delete(value)**: Remove a node while preserving the BST ordering rule. Handle leaf nodes, nodes with one child, and nodes with two children separately.
- **InOrder()**: Visit left subtree, current node, then right subtree. For a BST, this produces sorted values.
- **PreOrder()**: Visit current node, then left subtree, then right subtree. Useful for copying or serializing tree shape.
- **PostOrder()**: Visit left subtree, then right subtree, then current node. Useful when deleting or freeing nodes bottom-up.

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

### **Method Notes**

- **Push(value)**: Add the value at the end of the backing array, then repeatedly swap it with its parent until the heap invariant is restored.
- **Pop()**: Remove and return the root value. Move the last value to the root, shrink the array, then repeatedly swap it with the smaller child until the heap invariant is restored.
- **Peek()**: Return the root value without removing it. This is the minimum value in a min-heap.
- **HeapifyUp(index)**: Move a newly inserted value upward while it has higher priority than its parent.
- **HeapifyDown(index)**: Move a displaced value downward while a child has higher priority.

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

### **Method Notes**

- **Put(key, value)**: Hash the key to choose a bucket. Update the value if the key already exists, otherwise insert a new entry and resize if the load factor gets too high.
- **Get(key)**: Hash the key, search the matching bucket, and return the stored value if the key exists.
- **Delete(key)**: Hash the key, remove the matching entry from its bucket, and update the item count.
- **Resize()**: Allocate a larger bucket array and reinsert every existing key-value pair so each entry is placed according to the new capacity.
- **Hash(key)**: Convert a key into a stable integer bucket index. For practice, start with string keys before trying generic keys.

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

### **Algorithm Notes**

- **BubbleSort(values)**: Repeatedly compare neighboring values and swap them when they are out of order. Each pass moves one large value toward the end.
- **InsertionSort(values)**: Walk left to right and insert each value into the already-sorted portion before it.
- **MergeSort(values)**: Split the input into halves, sort each half recursively, then merge the sorted halves into a new sorted result.
- **QuickSort(values)**: Choose a pivot, partition values around it, then recursively sort the partitions.
- **HeapSort(values)**: Build a heap from the values, then repeatedly remove the top-priority value to produce sorted output.

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

### **Method Notes**

- **AddVertex(value)**: Add a vertex to the adjacency list if it does not already exist.
- **AddEdge(from, to)**: Add `to` to `from`'s neighbor list. For an undirected graph, also add `from` to `to`'s neighbor list.
- **RemoveEdge(from, to)**: Remove the neighbor relationship between two vertices. For an undirected graph, remove both directions.
- **RemoveVertex(value)**: Delete the vertex and remove it from every other vertex's neighbor list.
- **BFS(start)**: Visit vertices level by level using a queue and a visited set.
- **DFS(start)**: Visit as far as possible along each path using recursion or an explicit stack plus a visited set.

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

### **Method Notes**

- **Get(key)**: Look up the node in the hash map. If found, move it to the front of the linked list and return its value.
- **Put(key, value)**: Update an existing node or create a new one at the front. If capacity is exceeded, evict the tail node and remove its key from the map.
- **MoveToFront(node)**: Unlink a node from its current position and reinsert it at the head to mark it as recently used.
- **EvictLeastRecent()**: Remove the tail node because it represents the least recently used item.

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
