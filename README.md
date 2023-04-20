# Card Deck Service: Go & SQLite Implementation

A versatile Card Deck Service implemented in Go and SQLite, designed to be used in card games such as Poker, Bridge, 29, etc.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Ubuntu](#ubuntu)
  - [macOS](#macos)
- [Build and Run](#build-and-run)
  - [Build](#build)
  - [Run](#run)
- [Run Tests](#run-tests)
- [Engineering Design Choices](#engineering-design-choices)
  - [Functional Requirements](#functional-requirements)
  - [Non-Functional Requirements](#non-functional-requirements)
  - [High-Level Design](#high-level-design)
  - [Component Design](#component-design)
  - [Database Design](#database-design)
  - [Database Choice](#database-choice)

## Prerequisites

Ensure GoLang is installed on your system.

## Installation

### Ubuntu

{Fill the required steps}

### macOS

{Fill in the required steps using brew incl. brew install steps.}

## Build and Run

### Build

```sh
make build
```

### Run
```
make run
```

### Run Tests
```
make test
```

### Engineering Design Choices

## Functional Requirements

- Build a New Deck: Ability to build a new card deck that can be acted on.
- Open A Deck: Ability to open a specified card deck to show the remaining cards in the deck.
- Draw a Card: Ability to draw a specified number of cards from a specified deck

## Non-Functional Requirements

- Provide proper error messages to the user such as:
a. Deck does not exist
b. Number of cards to be drawn exceeds the number of cards present in the deck.
- Query Parameters to be presented through URL instead of request body.
- Abstraction Layers in Services & Data Access Layer to make it extensible to implement different Databases in the future.
- Proper Dependency Injection to each request context without initializing Objects for every request.

## High-Level Design

This is overall a simple API, so we would build a REST API in MVC Architecture. Separating concerns and responsibilities in different modules.
I have chosen to do a Monolithic System Architecture instead of Service Oriented or a Microservice Architecture since we would have services for only one model.
This can be turned into SOA or microservices easily if we put routes, tests, and internal folders to a subfolder called `deck`.

## Component Design

There's essentially only 1 component to be built which is Deck which consists of Cards.
As mentioned earlier we would separate out concerns.
Following submodules are built:

- Routing
- Services Abstraction
- Dependency Injection of correct Dependencies
- Settings separated out in another file to have different settings in different environments.

## Database Design

## Models

Deck
- id : string
- shuffled : boolean
- cards : [] Card

Card
- value
- suit
- code

Deck contains Card as a JSONB field.

We would not have a different table for Card but store it in a JSONB column in the Deck Table.

The pitfalls for having a different table for Card:

- Suppose we have a table Cards with 52 cards.
- For each deck, we would need to maintain a relationship table with FK in Cards & Deck.
- This junction table would Linearly grow with each new Deck.
      1 deck - 52 Relations
      10 Decks - 520 Relations.
      
These relations would need to be maintained i.e., Indices, records need to be deleted when cards are drawn, the order of the cards needs to be maintained in the relationship table as well.

This would be expensive in terms of Database Size + Operational Execution Time for the APIs.

### Database Choice

The database used in the project is SQLite because it's easier to run this on localhost. To productionize this, we can easily extend the DAL Layer for another Database.