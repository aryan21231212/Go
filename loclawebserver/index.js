import express from "express";
import cors from "cors";



const app = express();
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const PORT = 3000;


app.get("/",(req,res)=>{
    res.send("Welcome to the Noob's world");
})

app.get("/get",(req,res)=>{
    res.status(200).json({message:"Hello from Noob"});
})

app.post("/post",(req,res)=>{
    let myjson = req.body;
    res.status(200).send(myjson);
})


app.post("/postform",(req,res)=>{
    res.status(200).send(JSON.stringify(req.body));
})



app.listen(PORT,()=>{
    console.log(`Server is running on http://localhost:${PORT}`);
})