ALTER TABLE rsvps RENAME COLUMN nom TO name;
ALTER TABLE rsvps RENAME COLUMN presence TO attending;
ALTER TABLE rsvps RENAME COLUMN regime_alimentaire TO diet;
ALTER TABLE rsvps RENAME COLUMN participation_brunch TO brunch_participation;
ALTER TABLE rsvps RENAME COLUMN commentaire TO comment;

ALTER TABLE rsvps ADD COLUMN gite BOOLEAN DEFAULT FALSE NOT NULL;
ALTER TABLE rsvps ALTER COLUMN gite DROP DEFAULT;
ALTER TABLE rsvps ALTER COLUMN brunch_participation SET DEFAULT FALSE;
ALTER TABLE rsvps ALTER COLUMN brunch_participation SET NOT NULL;
ALTER TABLE rsvps ALTER COLUMN brunch_participation DROP DEFAULT;
