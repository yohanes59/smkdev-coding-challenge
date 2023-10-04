function checkCats(dataTuti, dataNining) {
	const koreksiDataTuti = dataTuti.slice(1, -2);
	const koreksiDataNining = dataNining.slice(1, -2);

	koreksiDataTuti.forEach((usia, index) => {
		if(usia >= 3) {
			console.log(`Data Tuti : Kucing Nomor ${index + 1} adalah Kucing Dewasa, dan berusia ${usia} tahun`);
		} else {
			console.log(`Data Tuti : Kucing Nomor ${index + 1} adalah Kucing anak-anak, dan berusia ${usia} tahun`);
		}
	});

	koreksiDataNining.forEach((usia, index) => {
		if(usia >= 3) {
			console.log(`Data Nining : Kucing Nomor ${index + 1} adalah Kucing Dewasa, dan berusia ${usia} tahun`);
		} else {
			console.log(`Data Nining : Kucing Nomor ${index + 1} adalah Kucing anak-anak, dan berusia ${usia} tahun`);
		}
	});
}

const dataTuti1 = [3, 5, 2, 12, 7];
const dataNining1 = [4, 1, 15, 8, 3];

const dataTuti2 = [9, 16, 6, 8, 3];
const dataNining2 = [10, 5, 6, 1, 4];

checkCats(dataTuti1, dataNining1);
console.log("\n");
checkCats(dataTuti2, dataNining2);