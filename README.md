# Rock Paper Scissors API

## Summary

This is an api that allows people to play Rock Paper Scissors remotely

## Build

This backend uses **Fiber**(_github.com/gofiber/fiber/v2_) to handle setup and manage api routes and **Gorm**(_gorm.io/gorm_) to manage interactions with the sqlite 3 database

## Routes

- Player Routes: **/rockpaper/player/all**[GET] -> **/rockpaper/player/leaderboard**[GET] -> **/rockpaper/player/add**[POST] -> **/rockpaper/player/getplayer**[POST] -> **/rockpaper/player/delete**[DELETE] -> **/rockpaper/player/tagline**[PUT]

- Game Routes: **/rockpaper/games/all**[GET] -> **/rockpaper/games/complete**[GET] -> **/rockpaper/games/incomplete**[GET] -> **/rockpaper/games/get**[POST] -> **/rockpaper/games/add**[POST] -> **/rockpaper/games/turn**[POST] -> **/rockpaper/games/cancel**[PUT] -> **/rockpaper/games/delete**[DELETE]

## Final Comments and Notes

- The api returns player objects that are currently broken(but do not impact API functionality)
- Wins do not increase when a player officially wins a game
