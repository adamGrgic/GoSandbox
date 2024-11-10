"use strict";
function init() {
    console.log('initialize test scoreboard module');
    console.log('setting score counts...');
    const player1Score = ['0'];
    const player2Score = ['0'];
    const player1Output = document.getElementById('player1_score_output');
    const player2Output = document.getElementById('player2_score_output');
    if (player1Output) {
        player1Output.innerHTML = player1Score?.at(-1) ?? 'NA';
    }
    if (player2Output) {
        player2Output.innerHTML = player1Score?.at(-1) ?? 'NA';
    }
    console.log('player scores have been set.');
    const player1MinusScoreBtn = document.getElementById('player1_minus_score_btn');
    player1MinusScoreBtn?.addEventListener("click", () => {
        console.log('player 1 minus btn clicked');
        const currentScore = player1Score?.at(-1) ?? 'NA';
        if (!currentScore) {
            console.error('player 1 current score not found');
            return;
        }
        const newScore = parseFloat(currentScore) - 1;
        player1Score.push(newScore.toString());
        if (player1Output) {
            player1Output.innerHTML = player1Score?.at(-1) ?? 'NA';
        }
    });
    const player2PlusBtn = document.getElementById('player2_plus_score_btn');
    player2PlusBtn?.addEventListener("click", () => {
        console.log('player 2 plus btn clicked');
        const currentScore = player2Score?.at(-1) ?? 'NA';
        if (!currentScore) {
            console.error('player 1 current score not found');
            return;
        }
        const newScore = parseFloat(currentScore) + 1;
        player2Score.push(newScore.toString());
        if (player2Output) {
            player2Output.innerHTML = player2Score?.at(-1) ?? 'NA';
        }
    });
    const player2MinusScoreBtn = document.getElementById('player2_minus_score_btn');
    player2MinusScoreBtn?.addEventListener("click", () => {
        console.log('player 2 minus btn clicked');
        const currentScore = player2Score?.at(-1) ?? 'NA';
        if (!currentScore) {
            console.error('player 1 current score not found');
            return;
        }
        const newScore = parseFloat(currentScore) - 1;
        player2Score.push(newScore.toString());
        if (player2Output) {
            player2Output.innerHTML = player2Score?.at(-1) ?? 'NA';
        }
    });
}
init();
//# sourceMappingURL=main.js.map