URL=http://localhost


ID=$(curl --silent ${URL}:8080/createLobby -d '{"playerId": "mat"}' | jq -r .'lobbyId')
echo Lobby ID: ${ID}
NEXT=$(curl --silent ${URL}:8080/joinLobby/${ID} -d '{"playerId": "pete"}' | jq -r .'board'.'nextToPlay')
echo Next To Play: ${NEXT}
echo Pete tries to play when it is not his turn...
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "pete", "tileLocation": 4}' | jq .'error'
echo Mat plays middle
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "mat", "tileLocation": 4}' | jq .'board'.'tiles'
echo Pete tries to play middle as well
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "pete", "tileLocation": 4}' | jq .'error'
echo Pete plays top middle
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "pete", "tileLocation": 2}' | jq .'board'.'tiles'
echo Mat plays middle right
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "mat", "tileLocation": 5}' | jq .'board'.'tiles'
echo Pete plays bottom right 
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "pete", "tileLocation": 8}' | jq .'board'.'tiles'
echo Mat plays middle left
curl --silent ${URL}:8080/makeMove/${ID} -d '{"playerId": "mat", "tileLocation": 3}' | jq