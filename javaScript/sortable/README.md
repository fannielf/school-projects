# Overview

The _Superhero Data Organizer_ is a web application designed to fetch, display, and manage data about various superheroes. The application allows users to view superhero statistics, sort the information, search by name, and paginate results effectively.

# Features

- **Data Fetching:** Retrieves superhero data from an external JSON API.
- **Dynamic Table:** Displays superhero data in a sortable and filterable table.
- **Pagination:** Supports multiple page sizes (10, 20, 50, 100, or all results) with a default of 20.
- **Search Functionality:** Interactive search feature that filters superheroes by name in real-time.
- **Sorting:** Users can sort the table by any column, toggling between ascending and descending order.

# Technologies Used

- HTML
- CSS
- JavaScript

# Getting Started

## Prerequisites

Before running this project, ensure you have:

- A code editor like [Visual Studio Code](https://code.visualstudio.com/)
- The [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) extension installed in VS Code

## Installation

1. Clone this repository:
   ```sh
   git clone https://01.gritlab.ax/git/fvesanen/sortable.git
   cd sortable
   ```
2. Open the project in **VS Code.**
3. Right-click on sortable.html and select **"Open with Live Server"**.

## Usage

1. Upon loading, the application fetches data from the superhero API.
2. The data is displayed in a table format with the following columns:

- Icon
- Name
- Full Name
- Powerstats (Intelligence, Strength, Speed, Durability, Power, Combat)
- Race
- Gender
- Height
- Weight
- Place of Birth
- Alignment

3. Use the pagination controls to navigate through the pages.
4. Select a page size from the dropdown to change the number of displayed results.
5. Type in the search bar to filter superheroes by name. The results update dynamically with each keystroke.
6. Click on table headers to sort the data by that column.

# Code Explanation

## Key Functions

- fetchData(): Fetches superhero data from the API and handles errors if the request fails.
- renderTable(data): Renders the data into an HTML table.
- sortAndRender(column, order): Sorts the data based on the specified column and order, then re-renders the table.
- convertWeightToKg(weight) and convertMeterToCm(height): Converts weight and height to a standard format for consistent sorting.
- getNestedValue(obj, path): Safely retrieves nested values from the superhero data object.

## Event Listeners

- Pagination controls to navigate between pages.
- Dropdown to change the number of displayed results.
- Search functionality that filters results in real-time.
- Click events on table headers to enable sorting.

# Authors

- [Fanni](https://github.com/fannielf)
- [Linnea](https://github.com/Linnie43)
- [Mahdi](https://github.com/mahdikheirkhah)
