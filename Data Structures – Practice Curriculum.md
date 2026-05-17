# **Data Structures – Practice Curriculum**

## **Overview**

This curriculum builds core data structure and algorithmic understanding from the ground up. It is **language agnostic**: examples are written in pseudocode, and the focus is on *concepts, invariants, and tradeoffs* rather than the syntax of any one language. Pick whatever language you are comfortable with (or learning) — the ideas transfer.

It assumes you can write basic code (variables, loops, functions, simple records/structs) and have used arrays/lists and maps as a consumer. It does **not** assume formal CS coursework.

Each section includes:

- **Background**: foundational ideas you need before reading the rest
- **Concept**: what the structure is and why it exists
- **Mental model / diagrams**: how to picture it
- **Implementation goals**: the methods to build
- **Method notes**: what each method does and the cases to handle
- **Worked example**: a small trace of the structure in action
- **Common pitfalls**: things that bite beginners
- **Test scenarios**: cases worth covering (great practice for a QA mindset)
- **Complexity**: cost of operations and why

---

## **0. Prerequisites: Concepts You Will Reuse Constantly**

Before starting, make sure these ideas are concrete. They appear in nearly every section.

### **Big-O Notation (Time and Space Complexity)**

Big-O describes how the cost of an operation grows as input size `n` grows. It is about *shape*, not absolute speed.

- **O(1)** — constant time. The cost does not depend on `n`. Example: reading `arr[5]`.
- **O(log n)** — logarithmic. Doubling `n` adds one more step. Example: binary search.
- **O(n)** — linear. Cost scales with input size. Example: scanning a list.
- **O(n log n)** — typical for good sorting algorithms.
- **O(n²)** — quadratic. Nested loops over the input. Example: bubble sort.

"Amortized O(1)" means *on average* across many operations the cost is constant, even if some individual operations are expensive. You will see this with dynamic arrays and hash maps.

### **References / Pointers**

A *reference* (or *pointer*) holds the location of a value in memory rather than the value itself. Languages express this differently — explicit pointers in C/C++/Go, references in Java/Python/JavaScript where every object variable is implicitly a reference — but the concept is universal.

Why references matter for data structures:
- Linked structures (lists, trees, graphs) connect nodes by reference.
- Methods that *modify* a structure must operate on the structure itself, not a copy of it.
- A "null" reference (`null`, `nil`, `None`, etc.) usually means "no node here" — accessing fields through a null reference is a runtime error.

### **Arrays vs Dynamic Arrays / Lists**

- **Fixed-size array**: a contiguous block of memory holding `N` elements. Size is set at allocation and cannot grow.
- **Dynamic array** (a.k.a. ArrayList, vector, list): wraps a fixed array but can "grow" by allocating a larger block, copying contents, and replacing the backing array. Most modern languages provide this as the default list type (`list` in Python, `ArrayList` in Java, `vector` in C++, `[]T` in Go, `Array` in JavaScript).

Two important properties:
- **Length**: how many elements are actually stored.
- **Capacity**: how many can be stored before another resize is needed.

You will reimplement this in Section 1 to see exactly how it works.

### **Records / Structs / Objects**

A grouped collection of named fields. Different languages call this different things (struct, class, record, object), but the concept is the same: a single value that holds multiple named fields.

```
record Node:
    value: integer
    next:  reference to Node
```

A linked structure is a graph of records connected through reference fields.

### **Recursion**

A function that calls itself. Each call gets its own local variables on the call stack. A recursive function needs:
1. A **base case** that returns without recursing.
2. A **recursive case** that makes progress toward the base case.

Recursion is natural for trees and divide-and-conquer algorithms. The cost is the call stack — too-deep recursion can stack-overflow.

### **Interfaces / Contracts**

Many languages let you define a *contract* (interface, abstract class, trait, protocol) that names methods without implementing them. Even without language support, it is good practice to define your data structure's API on paper before writing code:

```
Stack:
    push(value)
    pop() -> (value, ok)
    peek() -> (value, ok)
    length() -> integer
```

### **Generic Types (optional)**

Most modern languages support *generics*: a `Stack` of any type `T` rather than a `Stack` of `int`. For this curriculum, start with integers everywhere — it removes a layer of distraction. Add generics on a second pass if your language supports them.

### **A Note on Pseudocode**

Throughout this document, code is written in pseudocode that should look familiar regardless of language:

```
function append(value):
    if length == capacity:
        grow()
    backing[length] = value
    length = length + 1
```

Translate it into your chosen language as you go.

---

## **1. Dynamic Array (ArrayList / Vector / List)**

### **Background**

A plain array has a fixed size at creation. A dynamic array hides that limit by allocating a larger backing array and copying when needed. Most languages give you one out of the box — implementing it yourself shows you exactly how it works.

### **Concept**

A dynamic array maintains:

- **Length**: number of elements actually stored
- **Capacity**: size of the backing array

Invariant: `0 <= length <= capacity`.

When `length == capacity` and you try to add another element, you must **grow**: allocate a new backing array (typically 2× the old capacity), copy existing elements over, and replace the backing array. The new element goes into the now-available slot.

### **Mental Model**

```
backing: [ A, B, C, _, _, _, _, _ ]   capacity = 8
              length = 3
```

After `append(D)`:

```
backing: [ A, B, C, D, _, _, _, _ ]   capacity = 8, length = 4
```

After three more appends, length hits capacity. The next append triggers a grow:

```
old: [ A, B, C, D, E, F, G, H ]       capacity = 8, length = 8
                                      append(I) triggers grow
new: [ A, B, C, D, E, F, G, H, I, _, _, _, _, _, _, _ ]   capacity = 16, length = 9
```

### **Implementation Goals**

- `append(value)`
- `insert(index, value)`
- `remove(index)`
- `get(index)` / `set(index, value)`
- `length()` / `capacity()`

### **Method Notes**

- **append(value)**: If `length == capacity`, grow (allocate new backing of `2 * capacity`, or `1` if capacity is `0`, copy old contents). Then write the new value at index `length` and increment `length`.
- **insert(index, value)**: Validate `0 <= index <= length`. If `length == capacity`, grow. Shift elements `[index .. length-1]` one slot right (iterate from the right to avoid clobbering). Place the value at `index`, increment `length`.
- **remove(index)**: Validate `0 <= index < length`. Optionally save the removed value to return. Shift elements `[index+1 .. length-1]` one slot left. Clear the now-unused slot at `length-1` (if your language uses garbage collection and the array holds references, this lets the old value be collected). Decrement `length`.
- **get(index)**: Validate `0 <= index < length`. Return `backing[index]`.
- **set(index, value)**: Validate `0 <= index < length`. Assign `backing[index] = value`.
- **length()**: Return the stored length.
- **capacity()**: Return the backing array's size.

### **Worked Example**

```
Start:  capacity=2, length=0, backing=[_, _]
append(1):                    length=1, backing=[1, _]
append(2):                    length=2, backing=[1, 2]
append(3):  GROW to cap=4     length=3, backing=[1, 2, 3, _]
insert(0, 9):                 length=4, backing=[9, 1, 2, 3]
remove(2):  removed=2         length=3, backing=[9, 1, 3, _]
```

### **Common Pitfalls**

- **Off-by-one in insert**: index `length` is valid (it means "append at the end"); index `length+1` is not.
- **Shifting in the wrong direction**: when inserting, copy from right to left so you do not overwrite values you still need. When removing, copy from left to right.
- **Forgetting to grow**: writing into `backing[length]` when `length == capacity` is an out-of-bounds error.
- **Exposing the backing array**: if a method returns the underlying array directly, the caller can mutate your internals.

### **Test Scenarios**

- Append into an empty array (triggers initial grow from capacity 0 or 1).
- Append exactly to capacity, then once more (triggers grow).
- Insert at index 0, at `length`, and in the middle.
- Insert/get/set/remove at out-of-bounds indices (negative, `> length`).
- Remove from a single-element array.
- Many appends in a row to verify amortized cost.

### **Complexity**

| Operation       | Time             | Why                                    |
|-----------------|------------------|----------------------------------------|
| append          | Amortized O(1)   | Grows are rare; copy cost amortizes    |
| insert(i)       | O(n)             | Shifts elements right                  |
| remove(i)       | O(n)             | Shifts elements left                   |
| get / set       | O(1)             | Direct index                           |
| length / capacity | O(1)           | Stored as fields                       |

---

## **2. Singly Linked List**

### **Background**

A linked list is the first structure you build out of *nodes* connected by references. Unlike an array, the elements are not stored contiguously — each node lives wherever the allocator put it, and points to the next.

### **Concept**

A node holds a value and a reference to the next node. The list itself holds a reference to the head (and optionally the tail and a length counter).

```
head -> [1 | next] -> [2 | next] -> [3 | next] -> null
```

`null` marks the end of the list.

### **Mental Model**

Think of paper notes scattered across a table. Each note has a value and an arrow drawn to the next note. You only know where the chain starts (the head). To find the 5th note, you have to follow arrows from the head — there is no "give me note 5 directly".

### **Records**

```
record Node:
    value: integer
    next:  reference to Node (or null)

record List:
    head: reference to Node (or null)
    tail: reference to Node (or null)   // optional but makes pushBack O(1)
    size: integer                        // optional but makes length() O(1)
```

A null head means an empty list. Accessing `node.next` on a null `node` is an error — always check.

### **Pros / Cons**

**Pros**
- O(1) insertion at the head (and at the tail if you store a tail reference).
- No contiguous-memory requirement, no resize/copy cost.

**Cons**
- O(n) access by index — you must walk from the head.
- Extra memory per element (the `next` reference).
- Worse cache performance than arrays — nodes can be far apart in memory.

### **Implementation Goals**

- `pushFront(value)` / `pushBack(value)`
- `popFront()`
- `find(value)`
- `remove(value)`
- `length()`

### **Method Notes**

- **pushFront(value)**: Make a new node whose `next` is the current head. Set head to the new node. If the list was empty, also set tail to the new node.
- **pushBack(value)**: Make a new node. If tail is null (empty list), set both head and tail to it. Otherwise set `tail.next` to it and update tail. Without a tail reference, you must walk to the end first — that is O(n).
- **popFront()**: If head is null, return a "not found" signal (e.g., `(0, false)` or null). Save the head's value, advance head to `head.next`. If head is now null, also clear tail. Decrement size. Return the saved value.
- **find(value)**: Walk from head. Return the first node (or its index, depending on your design) whose value matches, or null/`-1` if none.
- **remove(value)**: Walk from head, tracking the previous node. When you find a match, set `prev.next = match.next`. Special case: removing the head means setting `head = head.next`. Update tail if you removed the last node.
- **length()**: Return the stored size, or walk and count if not stored.

### **Worked Example**

```
Start:           head=null, tail=null
pushBack(1):     head=[1]->null, tail=[1]
pushBack(2):     head=[1]->[2]->null, tail=[2]
pushFront(0):    head=[0]->[1]->[2]->null, tail=[2]
remove(1):       head=[0]->[2]->null, tail=[2]
popFront():      returns 0, head=[2]->null, tail=[2]
popFront():      returns 2, head=null, tail=null
popFront():      returns "not found" — empty list
```

### **Common Pitfalls**

- **Losing the rest of the list**: `head = newNode` *before* setting `newNode.next = head` orphans every node. Order of pointer assignments matters.
- **Forgetting to update tail**: removing or popping the last node leaves a dangling tail reference that points to a freed node.
- **Walking off the end**: stop when `current == null`, not `current.next == null`, unless you specifically need the second-to-last node (which you do for `remove`).
- **Stale size counter**: every mutation must update size, or `length()` lies.

### **Test Scenarios**

- All operations on an empty list.
- Single-element list: popFront, remove(only-value), remove(not-present).
- Remove from the head, the tail, and the middle.
- Find a value that does not exist.
- Push/pop alternation to verify references stay coherent.

### **Stretch Exercises**

- **Reverse in place**: walk the list, reversing each `next` reference. Use three references: `prev`, `curr`, `next`.
- **Cycle detection (Floyd's algorithm)**: two pointers, one moving 1 step at a time, the other 2. If they ever meet, there is a cycle.

### **Complexity**

| Operation     | With tail reference | Without tail   |
|---------------|---------------------|----------------|
| pushFront     | O(1)                | O(1)           |
| pushBack      | O(1)                | O(n)           |
| popFront      | O(1)                | O(1)           |
| find / remove | O(n)                | O(n)           |
| length        | O(1) if stored      | O(n) otherwise |

---

## **3. Doubly Linked List**

### **Background**

A singly linked list can only walk forward. If you want to remove a known node in O(1) — without re-walking from the head to find its predecessor — each node needs a reference back as well.

### **Concept**

Each node has `prev`, `value`, and `next`. The list maintains head and tail references.

```
null <- [prev | 1 | next] <-> [prev | 2 | next] <-> [prev | 3 | next] -> null
```

### **Mental Model**

A train where every car is coupled to the cars on both sides. Walking forward or backward is symmetric. Removing a car means uncoupling it from both neighbors and re-coupling those neighbors to each other.

### **Records**

```
record Node:
    value: integer
    prev:  reference to Node (or null)
    next:  reference to Node (or null)

record List:
    head: reference to Node (or null)
    tail: reference to Node (or null)
    size: integer
```

### **Implementation Goals**

- `pushFront(value)` / `pushBack(value)`
- `removeNode(node)` — O(1) given a node reference
- `traverseForward()` / `traverseBackward()`

### **Method Notes**

- **pushFront(value)**: Build new node with `prev=null`, `next=head`. If head exists, set `head.prev = newNode`. Set `head = newNode`. If tail was null, set `tail = newNode`.
- **pushBack(value)**: Mirror of pushFront. New node's `prev=tail`, `next=null`. If tail existed, set `tail.next = newNode`. Update tail (and head if list was empty).
- **removeNode(node)**: If `node.prev` is non-null, set `node.prev.next = node.next`. Else (node was head), set `head = node.next`. Similarly update `node.next.prev` or set `tail`. Decrement size. Optionally clear the removed node's references to help garbage collection.
- **traverseForward / traverseBackward**: Yield values from head→tail or tail→head following the appropriate reference.

### **Worked Example**

```
Start:                        head=null, tail=null
pushBack(1):  [1]              head=[1], tail=[1]
pushBack(2):  [1]<->[2]        head=[1], tail=[2]
pushFront(0): [0]<->[1]<->[2]  head=[0], tail=[2]
removeNode(node-1):
              [0]<->[2]        head=[0], tail=[2]
removeNode(head):
              [2]              head=[2], tail=[2]
removeNode(tail):
              empty            head=null, tail=null
```

### **Common Pitfalls**

- **Updating only one direction**: when linking or unlinking, both `prev` and `next` must be set. It is easy to forget the back-pointer.
- **Removing head/tail**: special cases — the neighbor on one side does not exist.
- **Removing a node not in this list**: there is nothing to detect this; the API trusts the caller. Document this assumption.

### **Test Scenarios**

- Forward and backward traversal on an empty list, single-element list, and multi-element list.
- removeNode at head, tail, middle.
- Push and remove repeatedly, verifying head/tail stay coherent.
- Remove the only node — both head and tail should become null.

### **Complexity**

| Operation         | Time |
|-------------------|------|
| pushFront/Back    | O(1) |
| removeNode(node)  | O(1) |
| traverseForward   | O(n) |
| find by value     | O(n) |

---

## **4. Stack and Queue**

### **Background**

Stacks and queues are not a new memory layout — they are *access patterns* enforced on top of an existing structure (array or linked list). The point is the discipline: a stack only lets you touch one end, a queue only lets you add at one end and remove at the other.

### **Stack (LIFO — Last In, First Out)**

The most recently pushed item is the next one popped. Like a stack of plates: you put plates on top and take them off the top.

Use cases:
- Function call frames (the call stack).
- Expression evaluation, matching brackets.
- Depth-first traversals (manual or recursive).
- Undo histories.

### **Queue (FIFO — First In, First Out)**

The oldest item is the next removed. Like a line at a checkout: first one in line is first one served.

Use cases:
- Breadth-first search (BFS).
- Task scheduling, job processing.
- Buffering between producer and consumer.

### **Backing Storage**

Both stacks and queues can be built on top of:
- A **dynamic array**: simple and fast for stacks. For queues, naïve `dequeue` from index 0 is O(n) (shifting elements). Two common fixes:
  - Use a **linked list** (head=front, tail=back): O(1) on both ends.
  - Use a **circular buffer (ring)**: an array with `head` and `tail` indices that wrap around modulo capacity. Resize when full.

### **Implementation Goals**

**Stack**
- `push(value)`
- `pop()` → `(value, ok)`
- `peek()` → `(value, ok)`
- `length()`

**Queue**
- `enqueue(value)`
- `dequeue()` → `(value, ok)`
- `peek()` → `(value, ok)`
- `length()`

### **Method Notes**

- **Stack push(value)**: Append the value to the backing storage.
- **Stack pop()**: If empty, return a "not found" signal. Otherwise read the last element, remove it, return it.
- **Stack peek()**: If empty, signal "not found". Otherwise return the last element without modifying the structure.
- **Queue enqueue(value)**: Append to the back.
- **Queue dequeue()**: If empty, signal "not found". Otherwise return the front. With an array and front-shifting, this is O(n) — accept that for the first pass and improve later. With a linked list or ring buffer it is O(1).
- **Queue peek()**: Return the front without removing it.

### **Worked Example (Stack)**

```
push 1:       [1]
push 2:       [1, 2]
push 3:       [1, 2, 3]
pop():    -> 3, [1, 2]
peek():   -> 2
pop():    -> 2, [1]
pop():    -> 1, []
pop():    -> "empty"
```

### **Worked Example (Queue)**

```
enqueue 1:    front=[1]
enqueue 2:    front=[1, 2]
enqueue 3:    front=[1, 2, 3]
dequeue(): -> 1, [2, 3]
peek():    -> 2
dequeue(): -> 2, [3]
dequeue(): -> 3, []
dequeue(): -> "empty"
```

### **Common Pitfalls**

- **Returning a default value silently when empty**: callers cannot distinguish "popped a real 0" from "stack was empty". Always return a status flag, an exception, or an "optional" wrapper.
- **Memory retention**: when shrinking, the old element may still be referenced by the backing storage. For arrays of references, explicitly clear the slot so the garbage collector can reclaim the value.
- **Queue O(n) dequeue**: a naïve "drop the first element" usually does not actually free memory in many implementations — the buffer just keeps growing.
- **Confusing direction**: in an array-backed queue, decide once whether the front is the lowest index or the highest index, and document it.

### **Test Scenarios**

- pop/dequeue from empty.
- peek on empty.
- push/pop and enqueue/dequeue alternation.
- Many enqueues followed by many dequeues (stress the backing storage).
- Single-element behavior.

### **Complexity**

Both should be O(1) for all operations. If your queue's dequeue is O(n), you have not actually built a queue yet.

---

## **5. Binary Search Tree (BST)**

### **Background**

Trees are a generalization of linked lists where each node can point to multiple children. A **binary** tree has at most two children per node (left and right). A **binary search tree** adds an ordering rule.

### **Concept**

For every node `N` in a BST:
- All values in `N`'s left subtree are **less than** `N.value`.
- All values in `N`'s right subtree are **greater than** `N.value`.

This rule is recursive — every subtree is itself a BST.

```
        8
       / \
      3   10
     / \    \
    1   6    14
       / \   /
      4   7 13
```

### **Mental Model**

Imagine binary search over a sorted array. At each step you compare to the middle, then go left or right. A BST is the same idea, but the structure itself encodes "go left or right". The shape depends on insertion order — different orders yield different trees.

### **Records**

```
record Node:
    value: integer
    left:  reference to Node (or null)
    right: reference to Node (or null)

record BST:
    root: reference to Node (or null)
```

Recursion is the natural way to write tree operations. Each function works on a node reference (a subtree's root) and returns the possibly-new root of that subtree.

### **Implementation Goals**

- `insert(value)`
- `search(value)` → `boolean`
- `delete(value)`
- Traversals: `inOrder`, `preOrder`, `postOrder`

### **Method Notes**

- **insert(value)**: Recurse from the root. If the current subtree is null, return a new leaf node. If the value is less than the current node, recurse left and assign the result to `node.left`. If greater, recurse right. (If equal, decide your policy: ignore duplicates, or always go right.)
- **search(value)**: Recurse similarly. If null, not found. If equal, found. Else recurse into the matching side.
- **delete(value)**: Three cases once you find the node:
    1. **Leaf** (no children): return null to the parent.
    2. **One child**: return that child to the parent.
    3. **Two children**: find the **in-order successor** (smallest value in the right subtree), copy its value into the current node, then recursively delete the successor from the right subtree.
- **inOrder(visit)**: Recurse left, visit current, recurse right. For a BST, this yields values in sorted order — a useful invariant for testing.
- **preOrder(visit)**: Visit current, recurse left, recurse right. Useful for serializing tree shape, since you can rebuild the tree from a pre-order sequence.
- **postOrder(visit)**: Recurse left, recurse right, visit current. Useful when you need to process children before the parent (e.g., freeing nodes).

### **Worked Example**

Insert sequence: 8, 3, 10, 1, 6, 14, 4, 7, 13

```
Insert 8:                Insert 3:                Insert 10:
   8                        8                        8
                           /                        / \
                          3                        3   10

Insert 1, 6:              Insert 14, 4, 7:         Final:
       8                       8                       8
      / \                     / \                     / \
     3   10                  3   10                  3   10
    / \                     / \    \               / \    \
   1   6                   1   6    14            1   6    14
                              / \                    / \   /
                             4   7                  4   7 13
```

In-order traversal of the final tree: `1, 3, 4, 6, 7, 8, 10, 13, 14` — sorted, as expected.

### **Common Pitfalls**

- **Forgetting to reassign the result**: `node.left = insert(node.left, v)` — if you call `insert` and ignore the return value, new nodes do not attach.
- **Delete with two children handled wrong**: copying just the successor's reference (instead of the value, then deleting the successor) can break the tree.
- **Unbalanced trees**: inserting sorted values (1, 2, 3, 4, ...) produces a degenerate right-leaning chain — a linked list with O(n) operations. A real BST library would self-balance (AVL, red-black). For practice, accept this and note the limitation.
- **Equality handling**: be deliberate about duplicates. Mixing policies leads to weird search behavior.

### **Test Scenarios**

- Insert, then search for present and absent values.
- In-order traversal must always return sorted output.
- Delete a leaf, a node with one child, a node with two children, the root.
- Delete from an empty tree (no-op).
- Insert sorted input — observe the degenerate shape; verify operations still work, just slowly.

### **Complexity**

| Operation | Average    | Worst (degenerate) |
|-----------|------------|--------------------|
| Insert    | O(log n)   | O(n)               |
| Search    | O(log n)   | O(n)               |
| Delete    | O(log n)   | O(n)               |
| Traversal | O(n)       | O(n)               |

The "average" assumes random insertion order. Without rebalancing, adversarial input degrades to O(n).

---

## **6. Heap (Priority Queue)**

### **Background**

A **priority queue** is a queue where you do not get items out in insertion order — you get them out in priority order. The most common implementation is a *heap*.

### **Concept**

A **binary heap** is a complete binary tree (every level full except possibly the last, which fills left to right) that satisfies the *heap property*:

- **Min-heap**: every parent ≤ its children. The minimum is at the root.
- **Max-heap**: every parent ≥ its children. The maximum is at the root.

Because the tree is always complete, you can store it in a flat array — no references needed.

### **Array Representation**

For a node at index `i`:
- Left child: `2i + 1`
- Right child: `2i + 2`
- Parent: `(i - 1) / 2` (integer division)

```
Tree:                      Array:
       1                   [ 1, 3, 2, 7, 5, 4, 8 ]
      / \                    0  1  2  3  4  5  6
     3   2
    / \  / \
   7  5 4   8
```

### **Mental Model**

Picture filling a swimming pool: water always seeks its lowest level. Inserting a value places it in the next empty leaf, then it "bubbles up" until the heap property holds. Removing the root replaces it with the last leaf, which then "sinks down" until it settles.

### **Implementation Goals**

- `push(value)`
- `pop()` → `(value, ok)`
- `peek()` → `(value, ok)`
- Internal: `heapifyUp(i)`, `heapifyDown(i)`

### **Method Notes**

- **push(value)**: Append to the end of the array. Call `heapifyUp(length-1)`.
- **pop()**: If empty, signal "not found". Save `data[0]`. Replace `data[0]` with the last element and shrink the array. If non-empty, call `heapifyDown(0)`. Return the saved value.
- **peek()**: Return `data[0]` without modifying.
- **heapifyUp(i)**: While `i > 0` and `data[parent(i)] > data[i]` (for a min-heap), swap them and set `i = parent(i)`.
- **heapifyDown(i)**: Loop. Find the smaller of `i`'s children (if any). If the smaller child is less than `data[i]`, swap and continue with `i` set to the child's index. Otherwise stop.

### **Worked Example (Min-Heap)**

```
push 5:    [5]
push 3:    [5, 3] -> heapifyUp -> [3, 5]
push 8:    [3, 5, 8] (already valid)
push 1:    [3, 5, 8, 1]
           heapifyUp(3): 1 < parent(5), swap -> [3, 1, 8, 5]
           heapifyUp(1): 1 < parent(3), swap -> [1, 3, 8, 5]
pop():     return 1
           replace root with last: [5, 3, 8]
           heapifyDown(0): smaller child is 3, 3 < 5, swap -> [3, 5, 8]
Result: returned 1, heap is [3, 5, 8]
```

### **Common Pitfalls**

- **Comparing only one child**: heapifyDown must pick the smaller child. Picking the left child unconditionally breaks the invariant.
- **Off-by-one in parent/child math**: parent of `0` is `0` in integer math — make sure your loop stops at `i == 0`.
- **Forgetting to shrink the array on pop**: you will keep "popping" the same stale element forever.
- **Unstable ordering**: heaps are not stable — equal-priority items can come out in any order.

### **Test Scenarios**

- pop on empty.
- Push values in increasing order, decreasing order, and random order — pop them all and verify they come out sorted.
- Repeated push/pop cycles.
- Verify the array satisfies the heap property after every operation.

### **Complexity**

| Operation | Time      |
|-----------|-----------|
| push      | O(log n)  |
| pop       | O(log n)  |
| peek      | O(1)      |
| Build from n values | O(n) (using bottom-up heapify) |

---

## **7. Hash Map**

### **Background**

A hash map gives you (on average) O(1) lookup of a value by an arbitrary key. It does this by turning each key into an array index via a *hash function*, then storing the key/value pair in a "bucket" at that index. When two keys hash to the same bucket — a *collision* — the bucket holds multiple entries.

### **Concept**

```
hash(key) -> integer -> integer mod numBuckets -> bucket index
```

A bucket is typically a list of `(key, value)` pairs. To find a key, you hash to its bucket, then linearly scan the bucket comparing keys.

The **load factor** is `numEntries / numBuckets`. As it rises, collisions multiply and operations slow down. At some threshold (commonly 0.75), the map *resizes*: doubles the bucket count and reinserts every entry under the new modulus.

### **Mental Model**

Imagine a wall of mailboxes. The hash function tells you which mailbox to use for a given name. Most mailboxes hold one letter, but if two names happen to hash to the same box, both letters end up there and you check the name on each envelope.

### **Records**

```
record Entry:
    key:   string
    value: any

record HashMap:
    buckets: array of (list of Entry)
    size:    integer
```

Start with **string keys** — string hashing is straightforward and avoids generics.

### **A Simple String Hash Function (FNV-style)**

```
function hash(s):
    h = 14695981039346656037   // FNV offset basis
    for each byte b in s:
        h = h XOR b
        h = h * 1099511628211   // FNV prime
    return h
```

For real production code, use your standard library's hash function. For learning, write your own to see how a simple stream of bytes becomes a well-distributed integer.

### **Implementation Goals**

- `put(key, value)`
- `get(key)` → `(value, ok)`
- `delete(key)`
- Internal: `resize()`, `hash(key)`

### **Method Notes**

- **put(key, value)**: Compute `hash(key) mod numBuckets`. Scan that bucket; if the key already exists, overwrite its value. Otherwise append a new entry, increment size. If `size / numBuckets > loadFactor`, call `resize()`.
- **get(key)**: Hash, scan the bucket linearly, return the value if found, else "not found".
- **delete(key)**: Hash, scan the bucket, remove the matching entry, decrement size.
- **resize()**: Allocate a new bucket array (typically `2 * numBuckets`). For every entry in every old bucket, recompute its bucket index against the new size and append to the new bucket. Replace the old buckets with the new.
- **hash(key)**: Whatever scheme you choose. For strings, FNV is simple and good enough for learning.

### **Worked Example**

Buckets array of size 4, simple hash returning `length(key) mod 4`.

```
put("a", 1):     hash=1, buckets[1] = [("a",1)]
put("bb", 2):    hash=2, buckets[2] = [("bb",2)]
put("ccc", 3):   hash=3, buckets[3] = [("ccc",3)]
put("dd", 4):    hash=2, buckets[2] = [("bb",2), ("dd",4)]   // collision
get("dd"):       hash=2, scan bucket[2], find "dd", return 4
delete("bb"):    hash=2, remove "bb", buckets[2] = [("dd",4)]
```

(Real hash functions distribute much better than `length(key) mod n`. This is just for illustration.)

### **Common Pitfalls**

- **Bad hash function**: `length(key) mod n` has terrible distribution. Even simple hashes like FNV are vastly better.
- **Forgetting to resize**: as load factor grows, every operation slows toward O(n).
- **Comparing hashes instead of keys**: two distinct keys can hash to the same value. Always compare the actual keys when scanning a bucket.
- **Mutating during iteration**: if your `get`/`delete` walks a bucket and modifies it mid-walk, you can skip entries.
- **Negative indices**: if you use a signed integer hash, `hash mod n` can be negative. Use unsigned types or `((hash mod n) + n) mod n`.

### **Test Scenarios**

- put then get the same key.
- put twice with the same key — get returns the latest value, size unchanged.
- get on a missing key.
- delete on a missing key (no-op).
- Insert enough entries to trigger a resize; get all of them afterward.
- Force collisions (use keys you know hash to the same bucket) and verify all are retrievable.

### **Complexity**

| Operation | Average | Worst |
|-----------|---------|-------|
| put       | O(1)    | O(n)  |
| get       | O(1)    | O(n)  |
| delete    | O(1)    | O(n)  |
| resize    | O(n) one-time, amortized O(1) per put |

Worst case happens when all keys collide in one bucket — usually a sign of a bad hash function or adversarial input.

---

## **8. Sorting Algorithms**

### **Background**

Sorting is a foundation: it makes searching, deduplication, and many algorithms easier. The classic algorithms are worth knowing not because you will write them in production (use your standard library) but because each illustrates a core algorithmic technique.

### **Definitions**

- **Stable**: a sort that preserves the relative order of equal elements. (Important when sorting records by one field.)
- **In-place**: uses O(1) or O(log n) extra memory beyond the input.
- **Comparison sort**: relies on comparing elements; no comparison sort can do better than O(n log n) in the general case.

### **Bubble Sort — O(n²)**

Repeatedly walk the array, swapping adjacent out-of-order pairs. Each pass "bubbles" the largest remaining unsorted element to its final position at the end.

Useful only as a teaching tool. The one virtue: very simple to implement and understand.

```
[5, 1, 4, 2, 8]
pass 1: [1, 4, 2, 5, 8]
pass 2: [1, 2, 4, 5, 8]
done
```

### **Insertion Sort — O(n²) worst, O(n) on nearly-sorted input**

Walk left to right. For each element, shift it leftward into its correct place within the already-sorted prefix. Like sorting a hand of playing cards.

Excellent for small arrays (n < ~20) and nearly-sorted data. Often used as the base case inside hybrid algorithms (e.g., introsort switches to insertion sort for small partitions).

### **Merge Sort — O(n log n), stable, O(n) extra memory**

Divide and conquer:
1. Split the array in half.
2. Recursively sort each half.
3. Merge the two sorted halves into one sorted array.

The merge step is the heart of the algorithm: with two pointers (one into each half), repeatedly take the smaller front element until both are exhausted.

```
mergeSort([5,1,4,2,8,3])
  mergeSort([5,1,4]) -> [1,4,5]
  mergeSort([2,8,3]) -> [2,3,8]
  merge:               [1,2,3,4,5,8]
```

Stable, predictable performance, but allocates extra arrays.

### **Quick Sort — O(n log n) average, O(n²) worst, in-place**

Divide and conquer:
1. Choose a *pivot* (often the last element, the middle, or a random element).
2. Partition the array so values less than the pivot come before it, values greater after.
3. Recurse on each side.

Faster than merge sort in practice (better cache behavior, no extra allocations) but the worst case happens with bad pivot choices on already-sorted input. Random or median-of-three pivots avoid this.

### **Heap Sort — O(n log n), in-place, not stable**

1. Build a max-heap from the input.
2. Repeatedly swap the root (the maximum) with the last unsorted position, shrink the heap by one, and heapify down.

Uses the heap structure from Section 6. Predictable O(n log n) and no extra memory, but slower in practice than well-implemented quicksort.

### **Algorithm Notes**

- **bubbleSort(values)**: Loop until a full pass makes no swaps.
- **insertionSort(values)**: For each `i` from 1 to length-1, shift `values[i]` left while it is smaller than its left neighbor.
- **mergeSort(values)**: Recursively split until length ≤ 1, then merge upward.
- **quickSort(values)**: Choose a pivot, partition, recurse on each side. Recurse on the smaller side first if you want bounded stack depth.
- **heapSort(values)**: Build heap, extract max repeatedly into the back of the array.

### **Common Pitfalls**

- **Quicksort on sorted input with a fixed last-element pivot**: O(n²). Randomize or use median-of-three.
- **Merge sort merge step**: forgetting to copy the remaining tail of one half after the other is exhausted leaves data behind.
- **Off-by-one in partition**: write a known-good partition once and reuse it.
- **Recursion depth**: very large inputs with poor pivot choices can blow the call stack.

### **Test Scenarios**

- Empty input, single-element input.
- Already sorted input, reverse-sorted input, all-equal input.
- Input with duplicates.
- Large random input — verify against your standard library's sort for cross-checking.
- For stable sorts: input of pairs `(key, originalIndex)` where keys repeat; verify original order is preserved.

### **Complexity Summary**

| Algorithm     | Best       | Average    | Worst      | Space     | Stable |
|---------------|-----------|------------|------------|-----------|--------|
| Bubble        | O(n)       | O(n²)      | O(n²)      | O(1)      | Yes    |
| Insertion     | O(n)       | O(n²)      | O(n²)      | O(1)      | Yes    |
| Merge         | O(n log n) | O(n log n) | O(n log n) | O(n)      | Yes    |
| Quick         | O(n log n) | O(n log n) | O(n²)      | O(log n)  | No     |
| Heap          | O(n log n) | O(n log n) | O(n log n) | O(1)      | No     |

---

## **9. Graphs**

### **Background**

A graph is the most general structure: a set of *vertices* (nodes) and *edges* (connections between them). Trees are graphs with extra rules; linked lists are graphs that happen to be a single path. Almost every relational problem — social networks, dependency resolution, shortest paths, routing — is a graph problem.

### **Concept**

- **Vertex**: a node with an identity (often a label or integer ID).
- **Edge**: a connection between two vertices. May be **directed** (one-way) or **undirected** (both directions).
- **Weighted graph**: edges carry a numeric weight (distance, cost, capacity).

```
Undirected:           Directed:
   A --- B               A --> B
   |     |               ^     |
   C --- D               |     v
                         C <-- D
```

### **Representations**

- **Adjacency list**: for each vertex, a list of its neighbors. Memory: O(V + E). Best for sparse graphs.
- **Adjacency matrix**: a `V x V` boolean (or weight) matrix where `M[i][j]` indicates an edge from `i` to `j`. Memory: O(V²). Fast edge lookup but wasteful for sparse graphs.

For this curriculum, use an adjacency list backed by a map:

```
record Graph:
    adjacency: map from vertex -> list of vertex
```

### **Mental Model**

Think of a city map. Vertices are intersections, edges are roads. Adjacency list = "for each intersection, here are the roads leading out." Adjacency matrix = "for every pair of intersections, is there a direct road?"

### **Implementation Goals**

- `addVertex(value)`
- `addEdge(from, to)` (undirected: add both directions)
- `removeEdge(from, to)`
- `removeVertex(value)`
- `bfs(start)` — traverse breadth-first
- `dfs(start)` — traverse depth-first

### **Traversals**

Both visit every reachable vertex from a starting point, but in different orders. Both need a **visited set** to avoid cycles.

#### **BFS — Breadth-First Search**

Uses a queue. Explores all vertices at distance 1, then all at distance 2, etc.

```
function bfs(start):
    visited = {start}
    queue = [start]
    while queue not empty:
        v = queue.dequeue()
        process(v)
        for n in neighbors(v):
            if n not in visited:
                visited.add(n)
                queue.enqueue(n)
```

Application: shortest path in unweighted graphs (the first time BFS visits a vertex is via the shortest path).

#### **DFS — Depth-First Search**

Uses recursion (or an explicit stack). Goes as deep as possible, then backtracks.

```
function dfs(v, visited):
    visited.add(v)
    process(v)
    for n in neighbors(v):
        if n not in visited:
            dfs(n, visited)
```

Application: cycle detection, topological sort, connectivity components.

### **Method Notes**

- **addVertex(value)**: If not already in the adjacency map, set its neighbor list to empty.
- **addEdge(from, to)**: Ensure both vertices exist (or error). Append `to` to `from`'s neighbors. For an undirected graph, also append `from` to `to`'s neighbors. Decide your policy on duplicate edges.
- **removeEdge(from, to)**: Filter out `to` from `from`'s neighbor list (and `from` from `to`'s if undirected).
- **removeVertex(value)**: Delete the vertex's entry. Walk every other vertex and filter `value` out of its neighbor list.
- **bfs(start)**: As pseudocode above; collect or yield each visited vertex.
- **dfs(start)**: As pseudocode above; recursive is the simplest implementation.

### **Worked Example**

Graph (undirected):
```
A - B
|   |
C - D - E
```

Adjacency list:
```
A: [B, C]
B: [A, D]
C: [A, D]
D: [B, C, E]
E: [D]
```

BFS from A: `A, B, C, D, E` (level by level).
DFS from A: `A, B, D, C, E` (depth-first; exact order depends on neighbor iteration order).

### **Common Pitfalls**

- **Forgetting the visited set**: graphs with cycles cause infinite loops.
- **Mutating the adjacency list during traversal**: copy or snapshot if you need to remove during iteration.
- **Asymmetric undirected edges**: addEdge must update both vertices' lists.
- **removeVertex leaves dangling references**: every other vertex's neighbor list must be cleaned up.
- **DFS recursion depth**: very deep graphs can blow the call stack. Use an explicit stack for large inputs.

### **Test Scenarios**

- Add and remove vertices and edges; verify the adjacency list state.
- BFS and DFS on a connected graph.
- BFS and DFS on a disconnected graph (only the start's component is reached).
- Graph with cycles — verify visited set prevents revisits.
- Self-loops (`A -> A`) and parallel edges if your design allows them.

### **Applications**

- BFS: shortest path in unweighted graphs, level-order tree traversal.
- DFS: cycle detection, topological sort, finding strongly connected components.
- Both: connectivity ("can I reach X from Y?").

### **Complexity**

For a graph with `V` vertices and `E` edges, using an adjacency list:

| Operation     | Time         |
|---------------|--------------|
| addVertex     | O(1)         |
| addEdge       | O(1)         |
| removeEdge    | O(degree)    |
| removeVertex  | O(V + E)     |
| BFS / DFS     | O(V + E)     |

---

## **10. LRU Cache (Capstone)**

### **Background**

A cache holds a fixed-size collection of recently used items. When it fills up, it must *evict* something to make room. **LRU (Least Recently Used)** is a popular policy: when full, evict the item that was used longest ago.

The challenge: support both `get(key)` and `put(key, value)` in O(1). Neither a list alone nor a map alone is enough — you need both together. This is why LRU is a classic interview problem and a perfect capstone: it forces you to *compose* data structures.

### **Concept**

Combine:
- **Hash map**: key → reference to the node. Gives O(1) lookup.
- **Doubly linked list**: holds `(key, value)` nodes in order from most-recently-used (head) to least-recently-used (tail). Gives O(1) reorder and O(1) eviction.

On every `get` or `put`, the affected node moves to the front. When capacity is exceeded, drop the tail.

### **Mental Model**

A bookshelf with a fixed number of slots. Every time you read a book, you put it back at the left end. When you buy a new book and the shelf is full, the rightmost book — the one nobody has touched in the longest — goes to the donation bin.

### **Why Both Structures?**

- A map alone tells you what is in the cache, but cannot answer "what is the least recently used?" in O(1).
- A list alone gives you ordering, but finding a key by name takes O(n).
- The map's value is a reference *into* the list. Lookup uses the map; reorder/eviction uses the list. Both are O(1).

### **Records**

```
record Entry:
    key:   any (hashable)
    value: any
    prev:  reference to Entry (or null)
    next:  reference to Entry (or null)

record LRUCache:
    capacity: integer
    size:     integer
    table:    map from key -> reference to Entry
    head:     reference to Entry (or null)   // most recently used
    tail:     reference to Entry (or null)   // least recently used
```

Storing the **key** inside the node is important — when you evict the tail, you need its key to delete the map entry.

### **Operations**

- **get(key)**: If present, move the node to head, return its value. Else return "not found".
- **put(key, value)**: If key exists, update its value and move its node to head. Else create a new node at head, add to map, increment size. If `size > capacity`, evict the tail node and remove its key from the map.

### **Method Notes**

- **get(key)**: Look up the node in the map. If missing, signal "not found". If found, call `moveToFront(node)`, return its value.
- **put(key, value)**: If key exists, update value, `moveToFront(node)`. Otherwise create node, insert at head, store in map, increment size. If `size > capacity`, call `evictLeastRecent()`.
- **moveToFront(node)**: Unlink the node from its current position (update its prev/next neighbors). Insert it at the head (update head's prev, the node's pointers, and head). If the node was the tail, update tail.
- **evictLeastRecent()**: Save the tail's key. Unlink the tail (set tail to its prev, update new tail's next to null, or set head to null if list becomes empty). Delete the saved key from the map. Decrement size.

### **Worked Example**

Capacity = 3.

```
put(1, "a"):     list: 1            map: {1}
put(2, "b"):     list: 2 -> 1       map: {1, 2}
put(3, "c"):     list: 3 -> 2 -> 1  map: {1, 2, 3}
get(1):    -> a, list: 1 -> 3 -> 2  map: {1, 2, 3}
put(4, "d"):     list: 4 -> 1 -> 3  map: {1, 3, 4}   // evicted 2
get(2):    -> not found
```

### **Common Pitfalls**

- **Forgetting to update the map on eviction**: the list shrinks but the map still holds the stale key, leaking memory.
- **moveToFront on the head node**: do nothing — but make sure your code handles it (don't unlink and relink incorrectly).
- **Tail update when removing the only node**: both head and tail go to null.
- **Updating value but forgetting to move to front**: a `put` on an existing key counts as a use. Move it.
- **Using the value (not the key) to look up in the map**: you need the key to evict from the map.

### **Test Scenarios**

- put up to capacity, then get all of them — none should be evicted.
- put past capacity — verify the right key is evicted and is no longer findable.
- get on a missing key.
- put with an existing key updates the value and moves to front.
- Sequence: put(1), put(2), put(3), get(1), put(4) — eviction should be 2 (LRU at the time of put(4)), not 1.
- Capacity 1 — every put evicts the previous entry.

### **Complexity**

All operations are O(1):

| Operation | Time |
|-----------|------|
| get       | O(1) |
| put       | O(1) |
| eviction  | O(1) |

### **Key Insight**

LRU is the canonical example of **composing** data structures to satisfy a constraint that no single structure can meet alone. This pattern — a map that points into another structure — recurs everywhere in systems work (database indexes, OS page caches, etc.).

---

## **Recommended Workflow**

For each structure:

1. **Write the interface first.** What methods does this structure expose? What types do they take and return? Do this on paper or as an interface/contract; resist jumping into implementation.
2. **Implement incrementally.** Get the simplest method (e.g., `length`, `push`) working first. Then add complexity.
3. **Add table-driven tests.** A loop that runs a set of `(input, expected)` rows is the most efficient way to cover many cases at once. As a QA-minded engineer this will feel natural.
   ```
   tests = [
       { name: "empty",    input: [],        expected: [] },
       { name: "single",   input: [1],       expected: [1] },
       { name: "sorted",   input: [1,2,3],   expected: [1,2,3] },
       { name: "reverse",  input: [3,2,1],   expected: [1,2,3] },
   ]
   for each tc in tests:
       run the case, assert output equals expected
   ```
4. **Add benchmarks** where it is interesting (e.g., `append` on the dynamic array, `push`/`pop` on the heap).
5. **Document Big-O complexity** in a comment near each method. This forces you to articulate *why* it has that complexity.
6. **Stress test**: run thousands of random operations and assert invariants (heap property, BST in-order is sorted, LRU size never exceeds capacity).

---

## **Final Notes**

- **Write without autocomplete or AI assistance for the first pass.** The point is to feel the syntax and structure, not produce working code fastest. Look up library docs, but type the code yourself.
- **Explain each structure out loud** as if teaching a junior engineer. If you cannot explain *why* a heap's parent index is `(i-1)/2`, you do not understand it yet.
- **Focus on tradeoffs, not just correctness.** Every structure exists because it is good at *something* and bad at *something else*. Write down both.
- **Reach for invariants when debugging.** "After every operation, the heap property holds." "After every put, the linked list is in MRU-to-LRU order." Bugs in data structures almost always violate an invariant.

This combination — building from the bottom up, testing rigorously, and articulating the *why* — will rebuild both coding fluency and deep CS intuition, regardless of the language you implement it in.

---

## **Appendix A: Writing Unit Tests in Go**

This curriculum is language-agnostic, but tests live in whatever language you pick. If you are using Go, the standard library's `testing` package is the only tool you need — no third-party assertion libraries, no test runners. This appendix is a high-level tour. Anything here can be cross-checked against the official docs at [go.dev/doc](https://go.dev/doc/) and the [`testing` package reference](https://pkg.go.dev/testing).

### **File Layout and Conventions**

- Tests live in files whose names end with `_test.go`. The Go toolchain only compiles these files when running `go test`; they are excluded from `go build`.
- Place test files in the **same directory** as the code they test. By convention they use the same package (so they can see unexported identifiers). Using `package foo_test` instead gives you a black-box test that can only touch the exported API — useful but rarely needed for this curriculum.
- A test function must be named `TestXxx` (the second letter must be upper-case) and take a single argument `t *testing.T`.

```go
package dynamicarray

import "testing"

func TestAppend(t *testing.T) {
    var a ArrayList
    a.Append(1)
    if got := a.Len(); got != 1 {
        t.Errorf("Len() = %d, want 1", got)
    }
}
```

### **Assertions: `t.Errorf` vs `t.Fatalf`**

Go has no `assert`. You compare values yourself with `if`, and report failures through methods on `*testing.T`:

- **`t.Errorf(format, args...)`** — record a failure but keep running the rest of the test. Use this for independent checks where later assertions still produce useful information.
- **`t.Fatalf(format, args...)`** — record a failure and stop the current test (or subtest) immediately. Use this when continuing would panic or produce noise (e.g., the value you wanted to inspect is `nil`).
- **`t.Helper()`** — call this at the top of a helper function so failure messages point at the *caller* rather than at the helper's line number. Essential for shared assertion helpers.

A good failure message includes the input, what you got, and what you wanted, in that order:

```go
t.Errorf("ReverseRunes(%q) = %q, want %q", in, got, want)
```

### **Running Tests**

From inside a module directory:

```sh
go test              # run all tests in the current package
go test ./...        # run all tests in this module recursively
go test -v           # verbose: print each test as it runs
go test -run TestAppend         # run only tests matching the regex
go test -run TestAppend/empty   # run a specific subtest (see below)
go test -cover       # report coverage percentage
go test -race        # enable the race detector (great for concurrent code)
go test -bench .     # run benchmarks (see below)
```

The `-run` flag takes a **regular expression** matched against the test name. `TestAppend/empty` targets a subtest named `empty` inside `TestAppend`.

### **Table-Driven Tests**

Table-driven tests are the canonical Go pattern for covering many input/output cases without duplicating boilerplate. The structure is always the same: a slice of anonymous structs, one row per case, and a loop.

```go
func TestReverse(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse(%q) = %q, want %q", c.in, got, c.want)
        }
    }
}
```

The pattern adapted from the official Go docs ([go.dev/doc/code](https://go.dev/doc/code)).

#### **Subtests with `t.Run`**

The simple loop above reports all failures under one test name, which gets noisy when many rows fail. Wrap each row in `t.Run` to give it its own name, isolate its failure, and make it individually runnable:

```go
func TestArrayList_Insert(t *testing.T) {
    cases := []struct {
        name    string
        start   []int
        index   int
        value   int
        want    []int
        wantErr bool
    }{
        {name: "empty/at zero",    start: []int{},        index: 0, value: 9, want: []int{9}},
        {name: "front",            start: []int{1, 2, 3}, index: 0, value: 9, want: []int{9, 1, 2, 3}},
        {name: "middle",           start: []int{1, 2, 3}, index: 1, value: 9, want: []int{1, 9, 2, 3}},
        {name: "end (== length)",  start: []int{1, 2, 3}, index: 3, value: 9, want: []int{1, 2, 3, 9}},
        {name: "out of bounds",    start: []int{1, 2, 3}, index: 4, value: 9, wantErr: true},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            a := newArrayList(tc.start)
            err := a.Insert(tc.value, tc.index)
            if (err != nil) != tc.wantErr {
                t.Fatalf("Insert error = %v, wantErr = %v", err, tc.wantErr)
            }
            if tc.wantErr {
                return
            }
            if got := a.Snapshot(); !slicesEqual(got, tc.want) {
                t.Errorf("after Insert(%d, %d): got %v, want %v", tc.value, tc.index, got, tc.want)
            }
        })
    }
}
```

Why this is worth the extra lines:

- Each row gets a stable name in `go test -v` output (`TestArrayList_Insert/middle`).
- You can re-run a single case with `go test -run TestArrayList_Insert/middle`.
- `t.Fatalf` inside a subtest stops *that subtest only*; the rest of the table still runs.
- Spaces in subtest names are converted to underscores when matched by `-run`.

#### **Subtest Gotchas**

- **Loop variable capture (pre-Go 1.22):** if you spell the loop as `for _, tc := range cases` and then close over `tc` inside `t.Run`, older Go versions share one variable across iterations. Modern Go (1.22+) makes the loop variable per-iteration, so this is no longer a foot-gun on `go 1.22+`. The modules in this repo use `go 1.26.2`, so you are safe.
- **`t.Parallel()`** marks a subtest as eligible to run in parallel with other parallel subtests. Useful for slow IO-bound tests, almost never needed for data-structure unit tests.

### **Benchmarks**

Benchmarks live next to tests, in `_test.go` files, and are named `BenchmarkXxx`. The framework runs the body `b.N` times and tunes `b.N` until it has a stable measurement.

```go
func BenchmarkArrayList_Append(b *testing.B) {
    var a ArrayList
    for i := 0; i < b.N; i++ {
        a.Append(i)
    }
}
```

Run with `go test -bench .`. Use `b.StopTimer()` / `b.StartTimer()` to exclude setup, and `b.ReportAllocs()` to include allocation counts in the output. Benchmarks are worth writing on structures where amortized vs worst-case timing matters (dynamic array growth, heap operations, hash map resize).

### **A Workable Default for This Curriculum**

For each data structure:

1. Define the public API (method signatures) before writing any test.
2. Start with one `TestXxx` per method that uses a table of subtests.
3. Always include rows for: empty input, single element, the bounds (`0`, `length-1`, `length`), out-of-bounds, and at least one "happy path" middle case.
4. Add invariant-style tests where they apply: heap property holds after every op, BST in-order is sorted, LRU size never exceeds capacity.
5. Benchmark only the operations whose complexity claims are non-obvious.

If a row fails, the failure message alone should tell you which input broke and what value was wrong — never make yourself re-read the test code to interpret a failure.
