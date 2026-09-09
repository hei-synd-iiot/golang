#import "@preview/grape-suite:4.0.0": exercise
#import exercise: project, task, hint, subtask

#let is-solution = "solution" in sys.inputs and sys.inputs.solution == "true"
#let hei-magenta = rgb(207, 20, 103)
#let hei-gray = rgb(127, 127, 127)
#let bonus = task.with(extra: true, numbering-format: (..n) => "")
#let task = task.with(numbering-format: (..n) => numbering("1", ..n))
#let subtask = subtask.with(markers: ("1.", "1."))
#let important = box.with(stroke: hei-magenta, fill: hei-magenta.lighten(95%), inset: 8pt, radius: 8pt)

#let sourcefile(filename, lang: "go") = {
    set text(
        font: "Jetbrains Mono",
        size: 7pt,
        weight: "light",
        fill: hei-gray
    )

    text[#filename]

    set text(
        weight: "regular",
    )
    let content = read(filename)
    sourcecode(raw(lang: lang, content))
}
#let shell(content) = {
    set text(
        font: "Jetbrains Mono",
        weight: "regular",
        fill: white
    )

    sourcecode(content, lang: "shell", numbering: none, frame: code-frame.with(stroke: hei-gray, fill: hei-gray.darken(50%)))
}


#show: project.with(
  no: 1,
  type: [Series],
  task-type: [Task],
  extra-task-type: [Bonus Task],
  show-outline: false,
  header-left: [#image("/docs/images/IIoT-Logo.svg", width: 24pt)],
  show-solutions: is-solution,

  title: [
    #text(font: "Avenir Next", weight: "extralight", fill: hei-gray, 16pt, [IIoT - Go Primer])

    #text(font: "Avenir Next", fill: hei-magenta, 32pt, [
      Mini Project
      #if is-solution [ Solution ]
    ])
      #image("/docs/images/golang.png", width: 100%)
  ],

  document-title: [
    #text(fill: hei-magenta, [*IIoT - Go Primer - Mini Project*])
  ],

  author: [_Rico Steiner, Michael Clausen_],

  show-hints: true,

  abstract: [
    You are building the software for a small monitoring station in a
    factory. The station polls *4 simulated sensors* (temperature probes, say) that
    each report a reading at random intervals. Over four stages, building on the
    Goroutines, Channels, Select, Timeouts, Closing Channels and WaitGroups/Mutexes
    sections of the course material, you will make the station:

    + print every reading as it arrives, tagged with which sensor it came from,
    + notice when a sensor stops reporting ("goes offline"),
    + run for a fixed duration and then shut down cleanly, printing a summary of how many readings it received from each sensor.

    Use the module's online _Go Prgramming Lnaguage Course_ (https://hei-synd-iiot.github.io/golang/) to make yourself familiar with Go's concepts used in this mini project.
  ],
)

#task(
  [One sensor, one channel],
  [Write a single simulated sensor and read its output.],
)[
  - Write a function `sensor(id int, out chan<- float64)` that, in a loop, sleeps
    for a random duration (e.g. between 300ms and 1200ms, using `time.Sleep` and
    `math/rand`) and then sends a random reading (e.g. a `float64` between 15.0
    and 30.0) on `out`.
  - Launch it with `go sensor(1, readings)` from `main`, and receive from
    `readings` in a loop in `main`, printing each value.
  - Let it run for a few seconds and confirm you see a steady, irregular stream
    of readings printed to the console.

  *Expected output* (values and timing will vary):
  ```
  sensor 1: 22.87
  sensor 1: 19.44
  sensor 1: 26.10
  ```
]

#task(
  [Fan-in multiple sensors with `select`],
  [Scale up to 4 sensors and merge their output.],
)[
  - Give each sensor its own channel, and start 4 sensor goroutines (ids 1-4).
  - In `main`, use a `select` statement with one `case` per sensor channel to
    receive whichever reading arrives first, and print it tagged with the
    sensor's id.
  - Wrap the `select` in a `for` loop so it keeps reacting to whichever sensor is
    ready next.

  #hint[
    If you find yourself copy-pasting the same print statement four times inside
    `select`, consider having every sensor send a small struct (id + value) on a
    single shared channel instead, and drop the `select` in favour of one
    `for range` loop. Both designs are valid --- try the `select` version first,
    it's what stage 3 builds on most directly.
  ]

  *Expected output*:
  ```
  sensor 2: 24.31
  sensor 1: 18.02
  sensor 4: 29.87
  sensor 3: 20.15
  ```
  Readings should arrive in no particular order, interleaved across sensors.
]

#task(
  [Detect an offline sensor],
  [Notice when a sensor stops reporting.],
)[
  - Decide on a deadline, e.g. 2 seconds: if a given sensor hasn't sent a reading
    within that time, the station should print e.g. `sensor 3: OFFLINE`.
  - Hint: for each sensor, `select` between receiving on its channel and
    `<-time.After(2 * time.Second)`. Each time you *do* receive a reading, the
    deadline effectively resets, because you call `time.After` again on the next
    loop iteration.
  - To test this, temporarily make one sensor goroutine stop sending after a few
    readings and confirm the dashboard correctly flags exactly that sensor as
    offline, while the other three keep reporting normally.

  #hint[
    A naive per-sensor `select` inside one shared loop only checks *one* sensor's
    timeout per iteration if you're not careful about structure. Make sure every
    sensor's timeout is tracked independently --- one working approach is one
    small goroutine *per sensor* that does its own `select` between "reading
    received" and "deadline expired", reporting to a single results channel that
    `main` prints from.
  ]

  *Expected output* (sensor 3 stopped early in this run):
  ```
  sensor 1: 21.44
  sensor 2: 25.03
  sensor 4: 18.76
  sensor 1: 20.91
  sensor 3: OFFLINE
  sensor 2: 24.55
  ```
]

#task(
  [Clean shutdown and a final summary],
  [Stop everything cleanly after a fixed run and report totals.],
)[
  - After a fixed duration (e.g. 15 seconds), the station should stop all sensor
    goroutines, wait for them to actually finish, and print a one-line summary
    per sensor: how many readings it received in total.
  - Use a shared `done chan struct{}` that `main` `close()`s when the run
    duration elapses. Every goroutine's loop should `select` on `done` alongside
    its normal work, and return when `done` is closed.
  - Use a `sync.WaitGroup` so `main` can block until every goroutine has actually
    returned before printing the summary --- don't just `time.Sleep` and hope.
  - Keep a per-sensor reading count. Since multiple goroutines update shared
    state, protect it with a `sync.Mutex` (or use one `atomic.Int64` counter per
    sensor if you'd rather avoid the mutex).

  #hint[
    Run your program with `go run -race .`. If you see a `DATA RACE` warning,
    some piece of state (probably your per-sensor counters) is being read or
    written from two goroutines without a mutex or atomic operation protecting
    it.
  ]

  *Expected output* (at the very end):
  ```
  shutting down...
  sensor 1: 14 readings
  sensor 2: 15 readings
  sensor 3: 3 readings
  sensor 4: 16 readings
  ```
]

#task(
  extra: true,
  [Stretch goal --- `context.Context`],
  [Replace the hand-rolled `done` channel with `context.WithTimeout`.],
)[
  Go's standard library provides a purpose-built way to express "run until
  cancelled or timed out": `context.Context`. Replace your `done` channel with a
  `context.Context`, and have each goroutine `select` on `ctx.Done()` instead.
  This isn't covered elsewhere in the course material, so treat it as an optional
  look ahead at a pattern you'll see constantly in real Go code, especially
  anything involving networking.
]
