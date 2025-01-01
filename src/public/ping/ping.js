

async function getPing() {
  const url = "http://localhost:6009/api/ping"
  try {
    const response = await fetch(url);
		console.log(response)
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const json = await response.json();
		return json

  } catch (error) {
    console.error(error.message);
  }
}

export { getPing }
