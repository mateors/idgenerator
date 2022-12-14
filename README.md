# idgenerator

> `go test -benchmem -run=^$ -bench ^Benchmark -benchtime=1s -count 1 idgenerator`

## Benchmark Output
![benchMarkOutput](./screens/benchmark_output.jpg)

## Winner Top to bottom
* [rs/xid](github.com/rs/xid) -> `Winner` |               = cdvo1etjdmi5bv2uetug
* [kjk/betterguid](github.com/kjk/betterguid)|            = -NHe2hg3DiBwZQF7tge9
* [chilts/sid](github.com/chilts/sid) |                   = 1669300411148099700-2359390368101562449
* [satori/go.uuid](github.com/satori/go.uuid) |           = b830ae52-1219-46c4-b98f-a5c072d8f7cd
* [google/uuid](github.com/google/uuid) |                 = 80ca72eb-8467-4767-95f1-4d0f0ab47dde
* [segmentio/ksuid](github.com/segmentio/ksuid) |         = 2HzvA4THd55V0fc4zpfmPT7hPed
* [oklog/ulid](github.com/oklog/ulid) |                   = 01GJN12V036S58T94RENX8QBB9
* [lithammer/shortuuid](github.com/lithammer/shortuuid) | = qqypVjXtN8qLx733jKnYEE
* [sony/sonyflake](github.com/sony/sonyflake) |           = 60c5b0ecd000067

<pre>
* github.com/rs/xid) -> `Winner`            = cdvo1etjdmi5bv2uetug
* github.com/kjk/betterguid                 = -NHe2hg3DiBwZQF7tge9
* github.com/chilts/sid                     = 1669300411148099700-2359390368101562449
* github.com/satori/go.uuid                 = b830ae52-1219-46c4-b98f-a5c072d8f7cd
* github.com/google/uuid                    = 80ca72eb-8467-4767-95f1-4d0f0ab47dde
* github.com/segmentio/ksuid                = 2HzvA4THd55V0fc4zpfmPT7hPed
* github.com/oklog/ulid                     = 01GJN12V036S58T94RENX8QBB9
* github.com/lithammer/shortuuid            = qqypVjXtN8qLx733jKnYEE
* github.com/sony/sonyflake                 = 60c5b0ecd000067
</pre>