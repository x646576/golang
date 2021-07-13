# CSP(Communicating Sequential Processes)

- Paper: [Communicating sequential processes](https://dl.acm.org/doi/10.1145/359576.359585) by C. A. R. Hoare
- Wikipedia: [Process calculus](https://en.wikipedia.org/wiki/Process_calculus)

## channel vs sync

| Question                                               | sync | channel |
| ------------------------------------------------------ | ---- | ------- |
| Is it a performance-critical section?                  | O    |         |
| Are you trying to transfer ownership of data?          |      | O       |
| Are you trying to guard internal state of a struct?    | O    |         |
| Are you trying to coordinate multiple pieces of logic? |      | O       |

## Go’s philosophy on concurrency

- Aim for simplicity
- Use channels when possible
- Treat goroutines like a free resource
