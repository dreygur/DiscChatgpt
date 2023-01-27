import { window } from "vscode";
import axios, { AxiosResponse } from 'axios';


export async function requestToOpenApi(query: string, apiKey: string): Promise<AxiosResponse> {
  let languageId = window.activeTextEditor?.document.languageId;

  /* eslint-disable @typescript-eslint/naming-convention */
  const options = {
    method: "POST",
    url: "https://api.openai.com/v1/completions",
    headers: {
      "Authorization": `Bearer ${apiKey}`,
      "content-type": "application/json",
    },
    data: {
      "model": "text-davinci-003",
      "prompt": `${query} (${languageId})`,
      "temperature": 0.7,
      "max_tokens": 256,
      "top_p": 1,
      "frequency_penalty": 0,
      "presence_penalty": 0,
    },
  };

  try {
    let response = await axios.request(options);
    if (response.status !== 200) { return response; }

    let text = response.data.choices[0].text;
    response.data = text.slice(2, text.length);
    return response;
  } catch (err: any) {
    // storeApiKey(undefined); // Remove the invalid API Key.
    return err.response;
  }
}