import axios from 'axios'

export async function getMailsxMatch(term, startTime, endTime, maxResults, fields){
    const config = {
        params: {
            term: term,
            startTime: startTime,
            endTime: endTime,
            maxResults: maxResults,
            fields: fields
        },
        headers: {
            Accept: "application/json",
        },
    };

    const url = "http://localhost:8080/search/match"
    try{
        console.log("esperando api...")
        const response = await axios.get(url, config);
        return await response.data;
    }
    catch (err){
        throw Error(err.response.statusText);
    }
   
}