# Escape-The-Maze
Finds the shortest path in a maze, using Dijkstra's algorithm

Description

The hero is trapped in a maze, there's mud up to his knees, and there are monsters in the maze! He must find a path so that he can safely escape!

Input:

Input is an ASCII-map of a maze. The map uses the following characters:

'#' for wall - The hero may not move here

' ' for empty space - The hero may move here, but only vertically or horizontally (not diagonally). Moving here costs the hero 1HP (health point) because of mud.

'm' for monster - The hero may move here, but only vertically or horizontally (not diagonally). Moving here costs the hero 11HP because of mud and a monster.

'S' this is where the hero is right now, the start.

'G' this is where the hero wishes to go, the goal. He can move here vertically or horizontally, costing 1HP.


Output:

The same as the input, but the route, which costs the least amount of HP is marked with '*', as well as the cost of the route.

Example

input:
```
######
#S  m#
#m## #
# m G#
######
```

output:
```
######
#S***#
#m##*#
# m G#
######
The path cost: 15 HP
```
