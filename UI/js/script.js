const X_CLASS = 'x'
const CIRCLE_CLASS = 'circle'

const cellElements = document.querySelectorAll('[data-cell]')
const board = document.getElementById('board')
const winningMessageElement = document.getElementById('winningMessage')
const startingMessageElement = document.getElementById('startingMessage')
const restartButton = document.getElementById('restartButton')
const xButton = document.getElementById('xButton')
const oButton = document.getElementById('oButton')
const winningMessageTextElement = document.querySelector('[data-winning-message-text]')

let player = 'X'
let user = null

xButton.addEventListener('click', () => {
    user = 'X';
    startGame();
})
oButton.addEventListener('click', () => {
    user = 'O';
    startGame();
    aiMove();
})
restartButton.addEventListener('click', () => {
    user = null;
    player = 'X'
    startGame();
})

startGame();

function startGame() {
    cellElements.forEach(cell => {
        cell.classList.remove(X_CLASS);
        cell.classList.remove(CIRCLE_CLASS);
        cell.addEventListener('click', handleClick);
    });
    setBoardHoverClass();
    if (user == null) {
        startingMessageElement.classList.add('show');
    } else {
        startingMessageElement.classList.remove('show');
    }
    winningMessageElement.classList.remove('show');
}

function makeMove(action, ai_move) {
    // Board --> server
    console.log(action)

    if (action.player !== player && !action.is_terminated) {
        return
    }

    if (action.move !== -1) {
        if (cellElements[action.move].classList.length > 1) {
            return
        }
        cellElements[action.move].classList.add(player === 'X' ? X_CLASS: CIRCLE_CLASS)
    }

    if (action.is_terminated) {
        if (action.winner === 'draw') {
            winningMessageTextElement.innerText = 'Draw!'
        } else {
            winningMessageTextElement.innerText = `${action.winner.toUpperCase()}'s Wins!`
        }
        winningMessageElement.classList.add('show')
        return
    }

    player = player === 'X' ? 'O' : 'X';

    setBoardHoverClass(player)

    if (ai_move) {
        aiMove();
    }

}

function aiMove() {
    let board = {}

    for (let i = 0; i < 9; i++) {
        if (cellElements[i].classList.length > 1) {
            board[`c${i}`] = cellElements[i].classList[1] === X_CLASS ? 'X' : 'O';
        } else {
            board[`c${i}`] = ''
        }
    }

    fetch("http://localhost:8080/updateBoard", {
        method: "POST", headers: { 'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify(board)
    }).then(async () => {
        let data = null
        fetch("http://localhost:8080/getMove").then(async rawResponse => {
            data = await rawResponse.json();
            // console.log(data);
            makeMove(data,false);
        });
    });
}

function handleClick(e) {
    const cell = e.target;

    if (player !== user) {
        return
    }

    const data = {is_terminated: false, winner: '', player: user, move: cell.dataset.cell};

    makeMove(data, true);
}

function setBoardHoverClass(player) {
    board.classList.remove(X_CLASS)
    board.classList.remove(CIRCLE_CLASS)
    if (player === 'O') {
        board.classList.add(CIRCLE_CLASS)
    } else {
        board.classList.add(X_CLASS)
    }
}
