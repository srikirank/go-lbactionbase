go-lbactionbase

## Go based library for Launchbar Model
[Launchbar](https://www.obdev.at/products/launchbar/index.html) for Mac is a productivity app that is very powerful and has a gazillion number of features. Some of them include:
* Launching Apps
* File Manager
* Clipboard Manager
* Web Searches
* Search Templates
* Command executors

The best feature is perhaps extensibility which helps extend the capabilities of Launchbar by letting developers create custom actions for their workflows.

This library has the following features:
1. Create LB suggestion items with helper methods.
2. Create error LB suggestion items for when there's an error while running the action
3. General file-based cache support for LB results with grace-time
4. Library methods for key LB directories: Actions, SharedScripts, Cache etc
5. Unix epoch to Human time converter to use in LB actions to populate subtitle, title, badge etc.