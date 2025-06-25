ALTER TABLE rsvps RENAME COLUMN name TO nom;
ALTER TABLE rsvps RENAME COLUMN attending TO presence;
ALTER TABLE rsvps RENAME COLUMN diet TO regime_alimentaire;
ALTER TABLE rsvps RENAME COLUMN brunch_participation TO participation_brunch;
ALTER TABLE rsvps RENAME COLUMN comment TO commentaire;

ALTER TABLE rsvps DROP COLUMN gite;
ALTER TABLE rsvps ALTER COLUMN brunch_participation DROP NOT NULL;