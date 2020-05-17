# ROLE DOC

_perrmission_table_
ver: 0.0.1

| operation         | admin | editor | author | contributer |
| ----------------- | ----- | ------ | ------ | ----------- |
| Post              |       |        |        |
| create            | Y     | Y      | Y      | Y           |
| update            | Y     | Y      | O      | O           |
| delete            | Y     | Y      | O      |             |
|                   |       |        |        |             |
| User              |
| createAdmin       | Y     |        |        |             |
| createEditor      | Y     |        |        |             |
| createAuthor      | Y     | Y      |        |             |
| createContributer | Y     | Y      | Y      |             |
| update            | Y     | O      | O      |             |
| delete            | Y     |        |        |             |
|                   |       |        |        |             |
| Setting           |       |        |        |             |
| create            | Y     | Y      | Y      |             |
| update            | Y     | O      | O      |             |
| delete            | Y     | O      | O      |             |
|                   |       |        |        |             |
| Tag               |       |        |        |             |
| create            | Y     | Y      |        |             |
| update            | Y     | O      |        |             |
| delete            | Y     | O      |        |             |
|                   |       |        |        |             |
| Comment           |       |        |        |             |
| delete            | Y     | Y      | O      |             |

> Y means Yes;
> N means No;
> O means only when the user is the content(mostly is post) creator
