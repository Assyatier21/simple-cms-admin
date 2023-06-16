#!/bin/bash

set -euo pipefail

echo "start migration"

MYSQL="mysql -h db -u root -proot"
DB_NAME="cms-repository"

$MYSQL $DB_NAME <<EOF
$(cat <<EOSQL
CREATE TABLE \`user\` (
  \`ID\` int(11) NOT NULL,
  \`Phone\` varchar(255) NOT NULL,
  \`Name\` varchar(255) NOT NULL,
  \`Role\` varchar(255) NOT NULL,
  \`Password\` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO \`user\` (\`ID\`, \`Phone\`, \`Name\`, \`Role\`, \`Password\`) VALUES
(1, '087786355690', 'Muhammad Sholeh', 'Admin', '\$2a\$08\$X/RsjSCkHNBMBNj/UGomQe/ezmrHtp1WqOl4wRrES0pnUvivCkd0e'),
(2, '09874273', 'Member Reguler', 'Member', '\$2a\$08\$X/RsjSCkHNBMBNj/UGomQe/ezmrHtp1WqOl4wRrES0pnUvivCkd0e');

ALTER TABLE \`user\`
  ADD PRIMARY KEY (\`ID\`);

ALTER TABLE \`user\`
  MODIFY \`ID\` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
EOSQL
)
EOF

echo "migration completed"

exec "$@"