# HelloSVC

This is an easy service to get started on, it simply publishes a message every
five seconds. It is also configurable over BW

To run this on your own machine, tweak the params.yml (or leave it as is) and
simply run it.

To run it on a spawnpoint, take a look at spawn.yml. To spawn it, do:
```
  spawnctl deploy -c spawn.yml -u scratch.ns
```
