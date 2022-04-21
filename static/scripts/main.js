new TypeIt("#welcome", {
  speed: 75,
  loop: false,
	breakLines: false,
})
.type("Welcome to my portflio.", {delay: 100})
.move(-4)
.type("o")
.move(4)
.go();

new TypeIt("#about", {
  strings: "  whoami && cat aboutme.txt",
  speed: 75,
  loop: false,
	breakLines: false,
}).go();

new TypeIt("#projects", {
  strings: "  ls ~/projects",
  speed: 75,
  loop: false,
	breakLines: false,
}).go();

new TypeIt("#links", {
  strings: "  cat ~/links.txt",
  speed: 75,
  loop: false,
	breakLines: false,
}).go();