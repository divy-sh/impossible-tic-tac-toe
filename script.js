const cells = document.querySelectorAll('.cell');
const statusString = document.querySelector('.status');
const restartButton = document.querySelector('.btn-restart');
let currentPlayer = 'X';
let board = ['', '', '', '', '', '', '', '', ''];
const winningCombinations = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6]
];

cells.forEach(cell => cell.addEventListener('click', handleCellClick));
restartButton.addEventListener('click', restartGame);

function handleCellClick(e) {
    const index = e.target.getAttribute('id');
    move(index);
    if (currentPlayer === 'O') { // AI move
        const bestMove = eval(board, 'O');
        if (bestMove !== null) {
            move(bestMove);
        }
    }
}

function move(index) {
    if (checkWin() || board.every(cell => cell !== '')) {
        return;
    }
    const element = document.getElementById(index);
    if (board[index] === '') {
        board[index] = currentPlayer;
        element.textContent = currentPlayer;
        if (checkWin()) {
            statusString.textContent = `Player ${currentPlayer} wins!`;
            cells.forEach(cell => cell.removeEventListener('click', handleCellClick));
        } else if (board.every(cell => cell !== '')) {
            statusString.textContent = `It's a draw!`;
        } else {
            currentPlayer = currentPlayer === 'X' ? 'O' : 'X';
            statusString.textContent = `Player ${currentPlayer}'s turn`;
        }
    }
}

function checkWin() {
    return winningCombinations.some(combination => {
        return combination.every(index => board[index] === currentPlayer);
    });
}

function restartGame() {
    board = ['', '', '', '', '', '', '', '', ''];
    cells.forEach(cell => {
        cell.textContent = '';
        cell.addEventListener('click', handleCellClick);
    });
    currentPlayer = 'X';
    statusString.textContent = `Player ${currentPlayer}'s turn`;
}

function eval(board, ply) {
    let bestScore = -Infinity;
    const moves = board.map((e, i) => (e === '' ? i : -1)).filter(i => i !== -1);
    if (moves.length === 0) {
        return null;
    }
    let bestMove = moves[0];
    for (let move of moves) {
        const newBoard = [...board];
        newBoard[move] = ply;
        const score = -negamax(newBoard, -Infinity, Infinity, false, ply === 'X' ? 'O' : 'X');
        if (score > bestScore) {
            bestScore = score;
            bestMove = move;
        }
        console.log(move, score);
    }
    return bestMove;
}

function negamax(board, alpha, beta, isMaximizingPlayer, ply) {
    if (winningCombinations.some(combination => {
        return combination.every(index => board[index] === 'X');
    })) {
        return ply === 'X' ? 1 : -1;
    } 
    if (winningCombinations.some(combination => {
        return combination.every(index => board[index] === 'O');
    })) {
        return ply === 'O' ? 1 : -1;
    }
    if (board.every(cell => cell !== '')) {
        return 0;
    }
    let bestScore = -Infinity;
    const moves = board.map((e, i) => (e === '' ? i : -1)).filter(i => i !== -1);
    for (let move of moves) {
        const newBoard = [...board];
        newBoard[move] = ply;
        const score = -negamax(newBoard, -beta, -alpha, !isMaximizingPlayer, ply === 'X' ? 'O' : 'X');
        bestScore = Math.max(bestScore, score);
        alpha = Math.max(alpha, score);
        if (alpha >= beta) {
            break;
        }
    }
    return bestScore;
}
