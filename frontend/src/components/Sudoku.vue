<template>
    <div class="content">
        <h1>{{ title }}</h1>
        
        <div>
            <form>
            <table class="game">
               
                <tr class="row" v-bind:key="row" v-for="(row, rowIndex) in sudoku">
                    <td class="cell" :class="{'border-right': colIndex === 2 || colIndex === 5, 
                                           'border-bottom': rowIndex === 2 || rowIndex === 5, 
                                           }"
                    v-bind:key="number" v-for="(number, colIndex) in row" >

                        <input name="" :value="number" :id="'row' + rowIndex + 'column' + colIndex" type="number" v-if="number != 0" readonly class="original">

                        <input name="" :id="'row' + rowIndex + 'column' + colIndex" type="number" min="1" max="9" maxlength="1" v-if="number === 0">
                    
                    </td>
                </tr>
                <br>  
            </table>
            <button type="submit" v-if="timerInit" class="verify">
                Resvisar
            </button>
            </form>
        </div>
        <br>
        <button @click="startNewGame" class="newmatch">
            Nueva partida
        </button>

    
         <div class="timer">
            <span class="elapsed">
                <span v-if="hours>0">{{ hours }}:</span>{{ minutes }}:{{ seconds }}
            </span>
            <button class="timer-btns" id="stop" v-if="timer" @click="stopTimer">
                <img src="../assets/images/pause.png">
            </button>
            <button class="timer-btns" id="start" v-if="showTimerVar && !timer" @click="showTimer">
                <img src="../assets/images/play-btn.png"> 
            </button>
            <br>
        </div>  
    </div>
        

   

</template>

<script>
    export default {
        name: "Sudoku",
        data() {
            return {
                title: "Sudoku Go",
                message: " ",
                sudoku: [],
                inputs: [],
                activeRow: -1,
                activeCol: -1,
                timer: null,
                timerInit:null,
                totalTime: 0,
                showTimerVar: null,
                name: ""
            };
        },
        methods: {
            getBoard: function() {
                 window.backend.generateSudoku().then(result => {
                 this.sudoku = result;
                 });
            },
            startNewGame: function() { 
                this.getBoard();
                this.stopTimer();
                this.resetTimer();
                this.showTimer();
            },
            showTimer(){
                this.timerInit=true
                this.showTimerVar=true
                this.startTimer();
            },
            resetTimer: function(){
                this.totalTime=0
            },
            setCellActive(row, col, original) {
                if (original) {
                    return
                }
                if(this.activeRow === row && this.activeCol === col) {
                    this.activeRow = -1
                    this.activeCol = -1
                    return
                }
                this.activeRow = row
                this.activeCol = col
            },

            startTimer: function() {
                if (!this.timer) {
                this.timer = setInterval(() => this.countdown(), 1000);
                }
            },
            countdown: function() {
                this.totalTime++;
            },
            stopTimer: function() {
                clearInterval(this.timer);
                this.timer = null;
                this.resetButton = true;
            },
            padTime: function(time) {
                return (time < 10 ? '0' : '') + time;
            },    
        },
        computed: {
            hours: function() {
                const hours = Math.floor(this.totalTime / (60*60))
                return hours;
            },          
            minutes: function() {
                const minutes = Math.floor(this.totalTime / 60) - (this.hours * 60);
                return this.padTime(minutes);
            },
            seconds: function() {
                const seconds = this.totalTime - (this.minutes * 60) - (this.hours *60*60);
                return this.padTime(seconds);   
            }
        }   
    }   

</script>

<style scoped>
    
    /* sudoku */
    .game {
        margin-left: 32%;
        width: calc(9*40px);
        margin: 0.5rem, auto;
    }
    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .cell {
        display: block;
        width: 40px;
        height: 40px;
        box-sizing: border-box;
        border: 1px solid rgb(58, 58, 58);
        background-color: #ccc;
        line-height: 40px;
        text-align: center;
    }
    .cell.border-right {
        border-right-width: 5px;    
    }
    .cell.border-bottom {
        border-bottom-width: 5px;    
    }
    .original {
        color: black;
        font-weight: bold;
        font-size:22px;
    }
    .cell:not(.original) {
        cursor: pointer;
    }
    form{
        align-content: center;
    }
    input {
        width: 32px;
        height: 35px;
        text-align: center;
        font-size: 20px;
        padding: 0;
        border: 3px transparent solid;
        background-color: transparent;
        color: rgb(75, 75, 75);
    }
    /* timer */
    .timer {
        font-size: 2em;
        color: #41BF98;
        text-align: center;
        padding:15px;
        margin:0 auto;
    }
    .timer-btns {
        border:none;
        color: #41BF98;
        font-size:1em;
        background-color: transparent;
        cursor: pointer;
    }
    .timer-btns img {
        width: 1em;
        height: 1em;
    }
    /*Botones */
    .newmatch {
	background:linear-gradient(to bottom, #404040 5%, #6b6b6b 100%);
	background-color:#404040;
	border-radius:18px;
	border:2px solid #41bf97;
	display:inline-block;
	cursor:pointer;
	color:#3fdfac;
	font-family: "Roboto", Helvetica, Arial, sans-serif;
	font-size:15px;
	padding:10px 40px;
	text-decoration:none;
    }
    .newmatch:hover {
	background:linear-gradient(to bottom, #6b6b6b 5%, #404040 100%);
	background-color:#6b6b6b;
    }
    .newmatch:active {
	position:relative;
    top:1px;
    }
    .verify {
    align-self: center;
	box-shadow: 0px 10px 14px -7px #276873;
	background:linear-gradient(to bottom, #41bf97 5%, #318f70 100%);
	background-color:#41bf97;
	border-radius:8px;
	display:inline-block;
	cursor:pointer;
	color:#ffffff;
	font-family:"Roboto", Helvetica, Arial, sans-serif;
    font-size:15px;
    height: 60px;
    width: 170px;
	text-decoration:none;
	text-shadow:0px 1px 0px #3d768a;
    }
    .verify:hover {
	background:linear-gradient(to bottom, #318f70 5%, #41bf97 100%);
	background-color:#318f70;
    }
    .verify:active {
	position:relative;
	top:1px;
    }
</style>