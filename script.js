const cells = document.querySelectorAll('.cell');
const statusString = document.querySelector('.status');
const restartButton = document.querySelector('.btn-restart');
const aiFirstToggle = document.getElementById('ai-first-toggle');
let currentPlayer = 'X';
let board = ['', '', '', '', '', '', '', '', ''];
const winningCombinations = [
    [0, 1, 2], [3, 4, 5], [6, 7, 8],
    [0, 3, 6], [1, 4, 7], [2, 5, 8],
    [0, 4, 8], [2, 4, 6]
];

cells.forEach(cell => cell.addEventListener('click', handleCellClick));
restartButton.addEventListener('click', restartGame);
aiFirstToggle.addEventListener('change', restartGame);

function handleCellClick(e) {
    const index = e.target.getAttribute('id');
    if (board[index] === '' && !checkWin() && !checkDraw()) {
        move(index);
        if (!checkWin() && !checkDraw()) {
            aiMove();
        }
    }
}

function move(index) {
    board[index] = currentPlayer;
    document.getElementById(index).textContent = currentPlayer;
    if (checkWin()) {
        statusString.textContent = `Player ${currentPlayer} wins!`;
        cells.forEach(cell => cell.removeEventListener('click', handleCellClick));
    } else if (checkDraw()) {
        statusString.textContent = `It's a draw!`;
    } else {
        currentPlayer = currentPlayer === 'X' ? 'O' : 'X';
        statusString.textContent = `Player ${currentPlayer}'s turn`;
    }
}

function aiMove() {
    const bestMove = eval(board, currentPlayer);
    if (bestMove !== null) {
        move(bestMove);
    }
}

function checkWin() {
    return winningCombinations.some(combination =>
        combination.every(index => board[index] === currentPlayer)
    );
}

function checkDraw() {
    return board.every(cell => cell !== '');
}

function restartGame() {
    board = ['', '', '', '', '', '', '', '', ''];
    cells.forEach(cell => {
        cell.textContent = '';
        cell.addEventListener('click', handleCellClick);
    });
    statusString.textContent = `Player ${currentPlayer}'s turn`;
    if (aiFirstToggle.checked) {
        aiMove();
    }
}

function eval(board, ply) {
    let bestScore = -Infinity;
    const moves = board.map((e, i) => (e === '' ? i : -1)).filter(i => i !== -1);
    if (moves.length === 0) return null;
    let bestMove = moves[0];
    for (let move of moves) {
        const newBoard = [...board];
        newBoard[move] = ply;
        const score = -negamax(newBoard, -Infinity, Infinity, ply === 'X' ? 'O' : 'X');
        if (score > bestScore) {
            bestScore = score;
            bestMove = move;
        }
    }
    return bestMove;
}

function negamax(board, alpha, beta, ply) {
    if (winningCombinations.some(combination => combination.every(index => board[index] === 'X'))) {
        return -1
    }
    if (winningCombinations.some(combination => combination.every(index => board[index] === 'O'))) {
        return -1
    }
    if (board.every(cell => cell !== '')) return 0;
    let bestScore = -Infinity;
    const moves = board.map((e, i) => (e === '' ? i : -1)).filter(i => i !== -1);
    for (let move of moves) {
        const newBoard = [...board];
        newBoard[move] = ply;
        const score = -negamax(newBoard, -beta, -alpha, ply === 'X' ? 'O' : 'X');
        bestScore = Math.max(bestScore, score);
        alpha = Math.max(alpha, score);
        if (alpha >= beta) break;
    }
    return bestScore;
}