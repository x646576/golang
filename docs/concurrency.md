# Concurrency

Why is concurrency hard?

## Race condition

Wikipedia: [Race condition](https://en.wikipedia.org/wiki/Race_condition)

**Race condition** arises in software when a computer program, to operate properly, depends on the sequence or timing of the program's processes or threads.

**Data race**: a situation where a memory operation in one thread could potentially attempt to access a memory location at the same time that a memory operation in another thread is writing to that memory location.

## Atomicity

Wikipedia: [Linearizability](https://en.wikipedia.org/wiki/Linearizability)

An atomic operation is one which cannot be (or is not) interrupted by concurrent operations.

## Critical section

Wikipedia: [Critical section](https://en.wikipedia.org/wiki/Critical_section)

**Critical section** or **critical region**: parts of the program where the shared resource is accessed need to be protected in ways that avoid the concurrent access.

## Memory Access Synchronization

Wikipedia:

- [Deadlock](https://en.wikipedia.org/wiki/Deadlock): a state in which each member of a group waits for another member, including itself, to take action, such as sending a message or more commonly releasing a lock.
- [Livelock](https://en.wikipedia.org/wiki/Deadlock#Livelock): similar to a deadlock, except that the states of the processes involved in the livelock constantly change with regard to one another, none progressing.
- [Starvation](<https://en.wikipedia.org/wiki/Starvation_(computer_science)>): a process is perpetually denied necessary resources to process its work.

### Necessary conditions

- Mutual exclusion: a non-shareable resource
- Hold and wait: resource holding
- No preemption: release by holder
- Circular wait: wait another process

## Determining Concurrency Safety

- Who is responsible for the concurrency?
- How is the problem space mapped onto concurrency primitives?
- Who is responsible for the synchronization?
