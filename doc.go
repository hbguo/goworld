/*

GoWorld

GoWorld is a distributed game server engine. GoWorld adopts a Space-Entity framework for game server programming.
Entities can migrate between spaces by calling `EnterSpace`. Entities can call each other using EntityID which is a
global unique identifier for each entity. Entites can be used to represent game objects like players, monsters, NPCs, etc.

Multiprocessing

GoWorld server contains multiple processes. There should be at least 3 processes: 1 dispatcher + 1 gate + 1 game.
The gate process is responsable for handling game client connections. Currently, gate supports multiple
transmission protocol including TCP, KCP or WebSocket. It also support data compression and encryption.
The game process is where game logic actually runs. A Space will always reside in one game process where it is created.
Entities can migrate between multiple game processes by entering spaces on other game processes.

*/
package goworld