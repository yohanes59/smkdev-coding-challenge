function generatePattern(number) {
	let pattern = "";
	for (let i = number; i >= 1; i--) {
		for (let j = i; j >= 1; j--) {
			pattern += j + " ";
		}
		pattern += "\n";
	}
	return pattern;
}

const patternInput = document.getElementById("pattern-input");
const generateButton = document.getElementById("generate-button");
const patternOutput = document.getElementById("pattern-output");

generateButton.addEventListener("click", () => {
	const inputValue = parseInt(patternInput.value);
	if (!isNaN(inputValue) && inputValue >= 1) {
		patternOutput.textContent = generatePattern(inputValue);
	} else {
		patternOutput.textContent = "Please enter a valid number.";
	}
});
