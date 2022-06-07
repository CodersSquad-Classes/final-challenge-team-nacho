Multithreaded Pacman Game - (single-node)
=========================================

Multithreaded version of the arcade video game [Pacman](https://en.wikipedia.org/wiki/Pac-Man). Features basic functionality to meet technical requirements. within scope and time constraints.
![Pacman](pacman.png)

Project Dependencies
-------
This project is built in GOlang 1.15, using the [Ebiten](https://ebiten.org/) game library for rendering, but running its own separate logic loop through goroutines (not the Ebiten Update loop). The project also makes use of GO's [image library](golang.org/x/image/font) for font resources.

Technical Functionalities
-------
- The game's maze layout is static.
- The `pacman` gamer is controlled by the user.
- Enemies are autonomous entities that will move in a random way.
- Enemies and pacman respect the layout limits and walls.
- Enemies number can be configured on game's start.
- Each enemy's behaviour is implemented as a separated thread.
- Enemies and pacman threads use the same map or game layout data structure resource.
- Displays obtained pacman's scores.
- Pacman loses a life when an enemy touches it and loses the game when dead 3 times.
- Pacman wins the game when it has taken all coins in the map.

Build/Run
-------
To build and run with default enemies (4), execute:

```bash
$ make run
```
To specify the number of enemies:

```bash
$ make run GHOSTNUM=i
```
> where i is the number of desired enemies in range [1, 4] (over or under will just adjust itself to the range)

Video Link
------
[VIDEO]()
