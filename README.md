- Requirements
  - GO
  - postgresql
 
- INSTALLATION
  - go install github.com/arturaciolii/gator
  - You will need to create a config file. Under your user's home directory create a .gatorconfig.json file.
    - Paste this inside it: {"db_url":"postgres://<user>:@localhost:5432/gator?sslmode=disable"}
