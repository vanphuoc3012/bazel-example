const express = require('express')

const Calculator = require('../node_calculator/calculator')

const app = new express()
const calculator = new Calculator()

app.listen(8080, () => {
    console.log("Start web app server with port 8080")
})

app.get('/',(_req, res) => {
    var num1 = getRandomArbitrary(1, 100)
    var num2 = getRandomArbitrary(1, 100)
    res.send(`did you know that ${num1} + ${num2} = ${calculator.add(num1, num2)}`)
})

function getRandomArbitrary(min, max) {
    return Math.random() * (max - min) + min;
  }