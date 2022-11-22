package org.pateason.leetcode;

import java.util.*;

/**
 * @see <a href="https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/description/>Leetcode</a>
 *
 * You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.')
 * and walls (represented as '+'). You are also given the entrance of the maze, where
 * entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are
 * initially standing at.
 *
 * In one step, you can move one cell up, down, left, or right. You cannot step into a
 * cell with a wall, and you cannot step outside the maze. Your goal is to find the nearest
 * exit from the entrance. An exit is defined as an empty cell that is at the border of
 * the maze. The entrance does not count as an exit.
 *
 * Return the number of steps in the shortest path from the entrance to the nearest exit,
 * or -1 if no such path exists.
 */
public class NearestExitFromEntranceInMaze {

  public static int nearestExit(char[][] maze, int[] entrance) {
    final char wallChar = '+';
    final int mazeHeight = maze.length;
    final int mazeWidth = maze[0].length;
    final int[][] directions = new int[][] {
        {-1, 0}, //north
        {0, 1}, // east
        {1, 0}, // south
        {0, -1} // west
    };

    final Queue<int[]> tileQueue = new ArrayDeque<>();
    tileQueue.add(entrance);

    int i = 0;
    int max = tileQueue.size();
    int steps = 0;
    while (!tileQueue.isEmpty()) {
      i++;
      if (i == max) {
        steps++;
        i = 0;
        max = tileQueue.size();
      }

      int[] currentTile = tileQueue.poll();
      int currentTileX = currentTile[1];
      int currentTileY = currentTile[0];

      // "close" current tile
      maze[currentTileY][currentTileX] = wallChar;

      // test directions
      for (int[] direction : directions) {
        int[] nextTile = new int[]{
            currentTileY + direction[0],
            currentTileX + direction[1]
        };

        // is at edge?
        if (
            (nextTile[0] < 0 || nextTile[0] == mazeHeight)
                || (nextTile[1] < 0 || nextTile[1] == mazeWidth)
        ) {
          continue;
        }

        // is a wall?
        if (maze[nextTile[0]][nextTile[1]] == wallChar) {
          continue;
        }

        // is an entrance?
        if (
            (nextTile[0] == 0 || nextTile[0] == mazeHeight - 1)
                || (nextTile[1] == 0 || nextTile[1] == mazeWidth - 1)
        ) {
          return steps;
        }

        if (tileQueue.stream().filter(x -> x[0] == nextTile[0] && x[1] == nextTile[1]).findFirst().isEmpty()) {
          tileQueue.offer(nextTile);
        }
      }
    }

    return -1;
  }
}
