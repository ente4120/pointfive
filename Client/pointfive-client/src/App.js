import "./App.css";
import React, { useEffect, useState } from "react"

function App() {
  const [events, setEvents] = useState([])
  const [actors, setActor] = useState([])
  const [repos, setRepos] = useState([])
  const [emails, setEmails] = useState([])

  const fetchEventsData = () => {
    fetch("http://localhost:8000/")
      .then(response => {
        return response.json()
      })
      .then(data => {
        setEvents(data.Events)
        setActor(data.Actors)
        setRepos(data.Repos)
        setEmails(data.Emails)
    })
  }

  useEffect(() => {
    fetchEventsData()
  }, [])

  return (
    <div className="App">
      <div className="App-header">Github Monitor ({events.length} Events Collected)</div>
      <div className="App-events-container">
        <div className="Events-container">
          <div className="App-sub-title">Events</div>
          <div className="Events-list">{events?.map(function(event) {
            return <div className="Event-container">
              <div className="Event-header">Event Id: <span>{event.Id}</span> Event type: <span>{event.Type}</span> Created: <span>{event.Created}</span></div>
              <div className="Event-body">Actor Name: <span>{event.Actor}</span> Actor URL: <a href={"https://github.com/"+event.Actor} target="_blank" >{"https://github.com/"+event.Actor}</a></div>
              <div className="Event-body">Repository Name: <span>{event.RepoUrl}</span> Repository URL: <a href={"https://github.com/"+event.RepoUrl} target="_blank">{"https://github.com/"+event.RepoUrl}</a></div>
            </div>
          })}</div>          
        </div>
        <div className="Repo-and-actor-container">
          <div>
            <div className="App-sub-title">Recent Repos</div>
            <div className="General-list">{repos.map(function(repo){
              return <div>
                <div className="Event-container">{repo}</div>
              </div>
            })}</div>
          </div>
          <div> 
          <div className="App-sub-title">Recent Actors</div>
          <div className="General-list">{actors.map(function(actor){
              return <div>
                <div className="Event-container">{actor}</div>
              </div>
            })}</div>
          </div>
          <div className="App-sub-title">Emails</div>
          <div className="General-list">{emails.map(function(mail){
              return <div>
                <div className="Event-container">{mail}</div>
              </div>
            })}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
