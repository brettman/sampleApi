# Sample API project
## This is a small api project, which deals with widget objects.  

I'm writing this as another golang demo.  I have several reasons for doing this...
- It's first weekend of the coronavirus quarantine in Kyiv, and I'm home bored
- I want to get comfortble with the standard go project layout and need practice
- Explore some common setup scenarios
    * Database setup
    * Dependency injection
    * API security
- Explore asynchronous programming in golang
    * _see UpdateCurrentValue on the WidgetService type_
- Explore devops considerations
    * deployment
    * testing
    * CI/CD

--- 
## Background and Caveats

I'm quite new to golang.  I will probably make a lot of mistakes. I have been working professionally with C# since the beta 2 release in... erm... 2001(??), so I'm not a total noob. I'm not uptight about learning, and welcome feedback, if you're so inclined. I've done other, similar projects which you may or may not find on my github account.  A deep forensic analysis of these projects will hopefully reveal a progression of skills and understanding, but really, who has time for all of that?  I hope that if you're still reading this, you might find it interesting...

Also, I'm writing this on Ubuntu, with open source tools: VsCode, VIM, etc.  I've been working on Microsoft windows professionally for more than 20 years, and this is really the first time I've decided to try to go whole hog with a linux project.  I have done a reasonable amount of small tasks with bash and know my VIM, but beyond that I'm pretty new to Linux.  I expect this will become apparent when I start trying to do deployment and all the devops stuff that always goes with a project.

You're probably thinking:  why bother with devops and deployment for a sample project?  Mostly, my experience on this is that code for code's sake is a total waste of time.  I'd much rather be surfing or spending time with my kids.  If I don't know how to deploy it, why would I bother writing it?

---
## Round One - Setup the basic project