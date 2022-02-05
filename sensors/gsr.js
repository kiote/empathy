const port = 22002;

function checkConnection() {
    console.log(`Trying to access NeuLog GSR sensor at http://localhost:${port}`);
}

export { checkConnection };