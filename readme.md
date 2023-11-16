# Connect Four Game

Welcome to the Connect Four game implemented in Go!

## Overview

This project is a text-based implementation of the classic Connect Four game. It supports two modes: playing against a computer opponent or playing with a friend.

The goal of this project was to take something that I was familiar with from Ruby, and use that as a vessel to learn more about Go!

## Features

- **Player vs. Computer:** Challenge the computer AI in a game of Connect Four.
- **Player vs. Player:** Play with a friend and take turns to connect four pieces in a row.

## Installation

1. Make sure you have [Go](https://golang.org/dl/) installed on your system.
2. Clone this repository:

    ```bash
    git clone https://github.com/sambcox/go-connect-four.git
    ```

3. Navigate to the project directory:

    ```bash
    cd go-connect-four
    ```

4. Run the game:

    ```bash
    go run main.go
    ```

## How to Play

1. Choose your game mode:
   - Press `c` to play against the computer.
   - Press `p` to play with a friend.
   - Press `q` to quit the game.

2. If playing with a friend, enter the names of Player 1 and Player 2.

3. During each turn:
   - Enter a letter between `A` and `G` to place your piece on the corresponding column.

4. The game continues until a player connects four pieces in a row or the board is full (resulting in a draw).

5. After the game ends, you can choose to play again or quit.

## Project Structure

The project is organized into the following packages:

- **board:** Contains the `Board` and `Cell` types, representing the game board and individual cells.
- **player:** Defines the `Player` type, representing a player in the game.
- **turn:** Manages the game turns, including user and computer moves.